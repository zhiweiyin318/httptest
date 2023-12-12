
1. build

```
go build -o bin/server ./server
go build -o  bin/client ./client
```

2. tls

```
openssl genrsa -out ca.key 2048

openssl req -new -x509 -days 36500 -key ca.key -out ca.crt -config openssl.cnf

openssl genrsa -out server.key 2048

openssl req -new -key server.key -out server.csr -config openssl.cnf

openssl x509 -req -days 36500 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -extfile <(printf "subjectAltName=IP.1:127.0.0.1")
```

2. run

```
go run server/server.go --cert="tls/server.crt" --key="tls/server.key"

go run client/client.go --addr="http://127.0.0.1:8090" --ca="tls/ca.crt"

curl --cacert "tls/ca.crt" -k https://localhost:8090/hello

```