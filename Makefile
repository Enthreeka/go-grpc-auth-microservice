

####### Protobuf #######
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/v1/auth.proto
	

####### Main #######
server:
	go run ./cmd/server/main.go

build:
	go build ./cmd/server/main.go


lint:
	echo "Starting linters"
	golangci-lint run ./... #--timeout=10m

test:
	go test -v ./internal/repo/postgres ./internal/repo/redis -parallel 2


####### Docker compose #######
dev:
	docker compose -f docker-compose.dev.yaml up
