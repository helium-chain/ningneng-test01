.PHONY: server
server: ./cmd/service/main.go ./tools/key/test.pem ./tools/key/test.key
	@go run ./cmd/service/main.go

.PHONY: client
client: ./cmd/client/main.go ./tools/key/test.pem
	@go run ./cmd/client/main.go

.PHONY: grpc
grpc: ./docs/proto/auth.proto
	@protoc -I ./docs/proto \
		--go_out ./internal/pb --go_opt paths=source_relative \
		--go-grpc_out ./internal/pb --go-grpc_opt paths=source_relative \
		auth.proto