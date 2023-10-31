#### Create Deposit Order Example Usage

run:
```bash
$ IP_ADDRESS=[PUT_SERVER_IP_ADDRESS_HERE] \
  PORT=[PUT_SERVER_PORT_HERE] \
  SERVER_CERT_PEM_PATH=./path/to/server_cert.pem \
  CLIENT_CERT_PEM_PATH=./path/to/client_cert.pem \
  CLIENT_KEY_PEM_PATH=./path/to/client_key.pem \
  go run examples/rpc-order-deposit-create/main.go
client_tx_id: tx7438041959
sum: 2570.50
currency_uuid: 4b96405b-654c-4f16-9527-eb42e650c8bb
```

output:
```bash
uuid: "bc29825c-5207-46c0-a0c2-cc9253bbf4ca"
status: "created"
```
