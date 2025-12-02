# ------------------------------
# Project Variables
# ------------------------------
PROTO_SRC=proto/llm.proto
PROTO_OUT=generated/proto

GO_BIN=$(shell go env GOPATH)/bin

# ------------------------------
# gRPC Code Generation
# ------------------------------
proto:
	@echo "🚀 Generating gRPC code..."
	@protoc \
		--go_out=generated/proto --go_opt=paths=source_relative \
		--go-grpc_out=generated/proto --go-grpc_opt=paths=source_relative \
	    proto/llm.proto
	@echo "✔ gRPC generation completed!"

# ------------------------------
# Run Go Server
# ------------------------------
run:
	@echo "🏃 Running Go WebSocket server..."
	@go run ./cmd/main.go

# ------------------------------
# Go Module Maintenance
# ------------------------------
tidy:
	@echo "🧹 Running go mod tidy..."
	@go mod tidy

# ------------------------------
# Build Binary
# ------------------------------
build:
	@echo "🔨 Building binary..."
	@go build -o bin/guessvibe ./cmd/main.go
	@echo "✔ Build completed! Binary: bin/guessvibe"

# ------------------------------
# Clean generated files
# ------------------------------
clean:
	@echo "🗑 Cleaning generated files..."
	@rm -rf $(PROTO_OUT)/*.pb.go
	@rm -rf bin/
	@echo "✔ Cleaned!"

# ------------------------------
# Help command
# ------------------------------
help:
	@echo "Available commands:"
	@echo "  make proto      - Generate gRPC + protobuf code"
	@echo "  make run        - Run WS server"
	@echo "  make tidy       - Clean go.mod/go.sum"
	@echo "  make build      - Build binary"
	@echo "  make clean      - Remove generated files"
