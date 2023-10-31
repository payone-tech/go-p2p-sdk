#### Fetch Banks Registry Example Usage

run:
```bash
$ IP_ADDRESS=[PUT_SERVER_IP_ADDRESS_HERE] \
  PORT=[PUT_SERVER_PORT_HERE] \
  SERVER_CERT_PEM_PATH=./path/to/server_cert.pem \
  CLIENT_CERT_PEM_PATH=./path/to/client_cert.pem \
  CLIENT_KEY_PEM_PATH=./path/to/client_key.pem \
  go run examples/rpc-banks-registry/main.go
```

output:
```bash
count: 2
uuid: "a8234838-db88-4476-b32b-49640899e685", title: "emirates bank"
uuid: "52c50471-a179-4c6d-9d53-393f5d35cb82", title: "capital bank"
```
