GOOGLEAPIS_DIR = ./google

generate-pb:
	protoc -I${GOOGLEAPIS_DIR} -I. --include_imports --include_source_info --descriptor_set_out=proto.pb ./hello.proto

generate-go:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        hello.proto

docker-run-envoy:
	docker run -it --rm --name envoy -p 8080:51051 \                                                                                              INT  21:52:37
      -v "$(pwd)/proto.pb:/protos/proto.pb:ro" \
      -v "$(pwd)/envoy.yaml:/etc/envoy/envoy.yaml:ro" \
      envoyproxy/envoy:v1.21.0
