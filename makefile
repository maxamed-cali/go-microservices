.PHONY: proto proto-product proto-order proto-category proto-common proto-user clean install-protoc-plugins run-all run-product run-order run-category run-api-gateway run-user

# Directory for generated proto files
PROTO_GEN_DIR = proto/gen

# Generate all proto files
proto: proto-common proto-product proto-order proto-category proto-user

# Generate product service proto
proto-product:
	mkdir -p $(PROTO_GEN_DIR)
	protoc --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative \
	       --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative \
	       proto/product.proto

# Generate common proto
proto-common:
	mkdir -p $(PROTO_GEN_DIR)
	protoc --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative \
	       --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative \
	       proto/common.proto

# Generate order proto
proto-order:
	mkdir -p $(PROTO_GEN_DIR)
	protoc --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative \
	       --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative \
	       proto/order.proto

# Generate category proto
proto-category:
	mkdir -p $(PROTO_GEN_DIR)
	protoc --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative \
	       --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative \
	       proto/category.proto

# Generate user proto
proto-user:
	mkdir -p $(PROTO_GEN_DIR)
	protoc --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative \
	       --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative \
	       proto/user.proto

# Clean generated files
clean:
	rm -rf $(PROTO_GEN_DIR)

# Install protoc plugins (run manually if needed)
install-protoc-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Run all services
run-all: run-product run-order run-category run-api-gateway run-user

# Run individual services
run-product:
	cd product && go run main.go

run-order:
	cd order && go run main.go

run-category:
	cd category && go run main.go

run-api-gateway:
	cd api-gateway && go run main.go

run-user:
	cd user && go run main.go
