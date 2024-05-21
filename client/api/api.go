package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/payone-tech/go-p2p-sdk/client/proto"
)

const (
	domain = "api.payone"

	MethodRegistryPaymentMethods = "client/registry/payment_methods"
	MethodRegistryCurrencies     = "client/registry/currencies"
	MethodOrderWithdrawCreate    = "client/order/withdraw/create"
	MethodOrderWithdrawCancel    = "client/order/withdraw/cancel"
	MethodOrderDepositCreate     = "client/order/deposit/create"
	MethodOrderDepositCancel     = "client/order/deposit/cancel"
	MethodOrderStatus            = "client/order/status"
	MethodReportToday            = "client/report/today"
	MethodReportYesterday        = "client/report/yesterday"
	MethodReport24hr             = "client/report/24hr"
)

// Client manages the RPC/SSE interface for a calling user.
type Client struct {
	ip             net.IP
	port           uint16
	serverCertPool *x509.CertPool
	clientCert     tls.Certificate
	httpClient     *http.Client
	rpcUrl         string
	sseEventsUrl   string
}

// MakeClient is the factory for constructing a Client for a given ip/port.
func MakeClient(
	ip net.IP,
	port uint16,
	serverCertPool *x509.CertPool,
	clientCert tls.Certificate,
	httpClient *http.Client,
) *Client {
	tlsCfg := &tls.Config{
		RootCAs:      serverCertPool,
		Certificates: []tls.Certificate{clientCert},
	}

	httpClient = makeHttpClient(httpClient, tlsCfg, ip)

	host := domain
	if port != 80 && port != 443 {
		host = host + ":" + strconv.FormatUint(uint64(port), 10)
	}
	url := &url.URL{Scheme: "https", Host: host}
	url.Path = "/rpc"
	rpcUrl := url.String()
	url.Path = "/sse/client/events"
	sseEventsUrl := url.String()

	c := &Client{
		ip:             ip,
		port:           port,
		serverCertPool: serverCertPool,
		clientCert:     clientCert,
		httpClient:     httpClient,
		rpcUrl:         rpcUrl,
		sseEventsUrl:   sseEventsUrl,
	}

	return c
}

// CallRpc performs a RPC request.
func (c *Client) CallRpc(
	rpcRequest proto.RpcRequest,
) (proto.RpcResponse, error) {
	var rpcResponse proto.RpcResponse

	b, err := json.Marshal(rpcRequest)
	if err != nil {
		return rpcResponse, err
	}

	req, err := http.NewRequest("POST", c.rpcUrl, bytes.NewReader(b))
	if err != nil {
		return rpcResponse, err
	}
	req.Header.Set("content-type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return rpcResponse, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("bad response status code: %d", resp.StatusCode)
		return rpcResponse, err
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return rpcResponse, err
	}

	err = json.Unmarshal(b, &rpcResponse)
	if err != nil {
		return rpcResponse, err
	}

	return rpcResponse, nil
}

// EventsStream performs a request to read server-side-events stream.
func (c *Client) EventsStream(
	ctx context.Context,
) (<-chan []byte, <-chan error, error) {
	req, err := http.NewRequest("GET", c.sseEventsUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("bad response status code: %d", resp.StatusCode)
		return nil, nil, err
	}

	ch := make(chan []byte)
	cherr := make(chan error)
	go func(ch chan<- []byte, cherr chan<- error) {
		var buf bytes.Buffer
		for {
			select {
			case <-ctx.Done():
				close(ch)
				close(cherr)
				return
			default:
			}

			b := make([]byte, 1024*4)
			n, err := resp.Body.Read(b)
			if err != nil {
				close(ch)
				cherr <- err
				close(cherr)
				return
			}
			buf.Write(b[:n])
			b = buf.Bytes()
			if len(b) < 2 {
				continue
			}
			if bytes.Compare(b[len(b)-2:], []byte("\n\n")) != 0 {
				continue
			}
			ch <- b
			buf.Reset()
		}
	}(ch, cherr)

	return ch, cherr, nil
}

func makeHttpClient(
	httpClient *http.Client,
	tlsCfg *tls.Config,
	ip net.IP,
) *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			// Timeout: 5 * time.Second,
		}
	}

	httpClient.Transport = &http.Transport{
		DialContext:            makeDialContext(ip),
		TLSClientConfig:        tlsCfg,
		TLSHandshakeTimeout:    5 * time.Second,
		DisableCompression:     true,
		MaxIdleConns:           2,
		MaxConnsPerHost:        2,
		IdleConnTimeout:        15 * time.Second,
		ResponseHeaderTimeout:  5 * time.Second,
		MaxResponseHeaderBytes: 1024 * 4,
		ForceAttemptHTTP2:      true,
	}

	return httpClient
}

type dialContext func(context.Context, string, string) (net.Conn, error)

func makeDialContext(ip net.IP) dialContext {
	var dialer net.Dialer

	dialContext := func(
		ctx context.Context,
		network,
		addr string,
	) (net.Conn, error) {
		addr = ip.String() + addr[strings.LastIndex(addr, ":"):]
		return dialer.DialContext(ctx, network, addr)
	}

	return dialContext
}
