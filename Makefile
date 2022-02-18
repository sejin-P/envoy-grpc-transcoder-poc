GOOGLEAPIS_DIR = ./google

generate-pb:
	protoc -I${GOOGLEAPIS_DIR} -I. --include_imports --include_source_info --descriptor_set_out=proto.pb ./hello.proto

generate-go:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        hello.proto
