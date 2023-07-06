proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/v1/auth.proto
	
server:
	go run ./cmd/server/main.go

build:
	go build ./cmd/server/main.go