package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/payone-tech/go-p2p-sdk/client/api"
	"github.com/payone-tech/go-p2p-sdk/client/proto"
	"github.com/payone-tech/go-p2p-sdk/examples/config"
)

func main() {
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

	request := proto.RpcRequest{
		Method: api.MethodReport24hr,
		Params: nil,
	}

	response, err := client.CallRpc(request)
	if err != nil {
		log.Fatalf("rpc call has failed: %v", err)
	}

	protoError := response.ProtoError()
	if protoError != nil {
		fmt.Println("error: %q (%d)", protoError.Message, protoError.Code)

		protoValidateErrors := protoError.ProtoValidation()
		for _, m := range protoValidateErrors {
			fmt.Errorf("field %q, reason %q", m.Field, m.Reason)
		}

		os.Exit(1)
	}

	var protoResponse *proto.ReportResponse
	err = json.Unmarshal(response.Result, &protoResponse)
	if err != nil {
		log.Fatal("unmarshal response result object has failed: %v", err)
	}

	fmt.Printf("total: %d\n", protoResponse.Total)
	fmt.Printf("total in: %d\n", protoResponse.TotalIn)
	fmt.Printf("total out: %d\n", protoResponse.TotalOut)
}
