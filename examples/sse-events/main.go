package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"github.com/payone-tech/go-p2p-sdk/client/api"
	"github.com/payone-tech/go-p2p-sdk/client/proto"
	"github.com/payone-tech/go-p2p-sdk/examples/config"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM)

	cfg := config.Params{}
	err := cfg.ReadEnv()
	if err != nil {
		log.Fatalf("load configuration has failed: %v", err)
	}

	client := api.MakeClient(
		cfg.IP,
		cfg.Port,
		cfg.ServerCertPool,
		cfg.ClientCert,
		(*http.Client)(nil))

	ch, cherr, err := client.EventsStream(ctx)
	if err != nil {
		log.Fatalf("open events stream has failed: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func(ch <-chan []byte, cherr <-chan error) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case b, ok := <-ch:
				if !ok {
					continue
				}
				event := parseEvent(b)
				fmt.Printf(
					"comment: %s\nevent: %s\ndata: %s\nretry: %s\n\n",
					string(event.Comment), string(event.Event),
					string(event.Data), string(event.Retry))
				if bytes.Compare(event.Event, []byte("order")) == 0 {
					var order *proto.OrderResponse
					err = json.Unmarshal(event.Data, &order)
					if err != nil {
						log.Fatal("unmarshal order object has failed: %v", err)
					}
					fmt.Printf("uuid: %q\n", order.UUID)
					fmt.Printf("status: %q\n", order.Status)
				}
			case err := <-cherr:
				fmt.Println(err)
				cancel()
				return
			default:
			}
		}
	}(ch, cherr)

	wg.Wait()
	cancel()
}

var (
	headerID      = []byte("id:")
	headerData    = []byte("data:")
	headerEvent   = []byte("event:")
	headerRetry   = []byte("retry:")
	headerComment = []byte(":")
)

func parseEvent(b []byte) proto.Event {
	fields := bytes.FieldsFunc(
		b,
		func(r rune) bool { return r == '\n' || r == '\r' })

	var event proto.Event
	for _, line := range fields {
		switch {
		case bytes.HasPrefix(line, headerComment):
			event.Comment = trimHeader(len(headerComment), line)
		case bytes.HasPrefix(line, headerID):
			event.ID = trimHeader(len(headerID), line)
		case bytes.HasPrefix(line, headerData):
			event.Data = append(event.Data, trimHeader(len(headerData), line)...)
			event.Data = append(event.Data, byte('\n'))
		case bytes.Equal(line, bytes.TrimSuffix(headerData, []byte(":"))):
			event.Data = append(event.Data, byte('\n'))
		case bytes.HasPrefix(line, headerEvent):
			event.Event = trimHeader(len(headerEvent), line)
		case bytes.HasPrefix(line, headerRetry):
			event.Retry = trimHeader(len(headerRetry), line)
		default:
		}
	}

	event.Data = bytes.TrimSuffix(event.Data, []byte("\n"))

	return event
}

func trimHeader(size int, b []byte) []byte {
	if b == nil || len(b) < size {
		return b
	}

	b = b[size:]
	if len(b) > 0 && b[0] == 32 {
		b = b[1:]
	}
	if len(b) > 0 && b[len(b)-1] == 10 {
		b = b[:len(b)-1]
	}

	return b
}
