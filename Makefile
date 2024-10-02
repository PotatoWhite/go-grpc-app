PROTO_SRC=proto/greeting.proto
PROTO_OUT=pkg/server/generated

all: build

build:
	@echo "Compiling Protobuf files..."
	protoc --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) $(PROTO_SRC)
	@echo "Building gRPC server..."
	go build -o bin/server cmd/go-grpc-app/main.go

clean:
	@echo "Cleaning generated files..."
	rm -rf $(PROTO_OUT)/*.go bin/server

run: build
	@echo "Running gRPC server..."
	./bin/server
