#### Fetch Currencies Registry Example Usage

run:
```bash
$ IP_ADDRESS=[PUT_SERVER_IP_ADDRESS_HERE] \
  PORT=[PUT_SERVER_PORT_HERE] \
  SERVER_CERT_PEM_PATH=./path/to/server_cert.pem \
  CLIENT_CERT_PEM_PATH=./path/to/client_cert.pem \
  CLIENT_KEY_PEM_PATH=./path/to/client_key.pem \
  go run examples/rpc-currencies-registry/main.go
```

output:
```bash
count: 1
uuid: "4b96405b-654c-4f16-9527-eb42e650c8bb", title: "uae dirham", code: "aed"
```
