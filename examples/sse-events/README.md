#### Subscribe on Order Events Example Usage

run:
```bash
$ IP_ADDRESS=[PUT_SERVER_IP_ADDRESS_HERE] \
  PORT=[PUT_SERVER_PORT_HERE] \
  SERVER_CERT_PEM_PATH=./path/to/server_cert.pem \
  CLIENT_CERT_PEM_PATH=./path/to/client_cert.pem \
  CLIENT_KEY_PEM_PATH=./path/to/client_key.pem \
  go run examples/sse-events/main.go
```

output:
```bash
comment: events stream
event:
data:
retry: 5000

comment:
event: order
data: {"uuid":"90dcb27e-cba9-4307-b889-ec81d8152d56","status":"created", …}
retry:

comment:
event: order
data: {"uuid":"90dcb27e-cba9-4307-b889-ec81d8152d56","status":"pending", …}
retry:

comment:
event: order
data: {"uuid":"90dcb27e-cba9-4307-b889-ec81d8152d56","status":"success", …}
retry:

…
```
