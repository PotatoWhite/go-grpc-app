### 

### Protobuf
#### Install
```bash
brew install proto
```

### Install (ubuntu)
```bash
sudo apt install protobuf-compiler
```

### Go Protobuf Plugin
```bash 
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Library
```bash
go get google.golang.org/grpc@v1.67.1
go get google.golang.org/protobuf@v1.34.2
```

## Makefile
```makefile
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
``` 
