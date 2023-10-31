#### Cancel Deposit Order Example Usage

run:
```bash
$ IP_ADDRESS=[PUT_SERVER_IP_ADDRESS_HERE] \
  PORT=[PUT_SERVER_PORT_HERE] \
  SERVER_CERT_PEM_PATH=./path/to/server_cert.pem \
  CLIENT_CERT_PEM_PATH=./path/to/client_cert.pem \
  CLIENT_KEY_PEM_PATH=./path/to/client_key.pem \
  go run examples/rpc-order-deposit-cancel/main.go
client_tx_id: tx7438041959
```

output:
```bash
status: "ok" (200)
```
