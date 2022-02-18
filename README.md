#RUN

go1.17, grpc, protoc 등의 환경 갖춰져있는지 확인


```go
go run cmd/main.go
```

```makefile
make docker-run-envoy
```

```shell
curl -X POST http://localhost:8080/say -H 'Conent-Type: application/json' -d '{
    "name": "abc",
    "hey": {
        "you": {
            "abc": "1"
        }
    }
}'
```

oneof 도 잘 proxy함을 확인.
