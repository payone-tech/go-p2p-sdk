#### Create Withdraw Order Example Usage

run:
```bash
$ IP_ADDRESS=[PUT_SERVER_IP_ADDRESS_HERE] \
  PORT=[PUT_SERVER_PORT_HERE] \
  SERVER_CERT_PEM_PATH=./path/to/server_cert.pem \
  CLIENT_CERT_PEM_PATH=./path/to/client_cert.pem \
  CLIENT_KEY_PEM_PATH=./path/to/client_key.pem \
  go run examples/rpc-order-withdraw-create/main.go
client_tx_id: tx7438041958
sum: 1570.50
account_number: 3030333305057070
currency_uuid: 4b96405b-654c-4f16-9527-eb42e650c8bb
```

output:
```bash
uuid: "0e810a3c-93ad-4472-bc05-8dd56d9c531f"
status: "created"
```
