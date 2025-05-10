# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Proto parameters
PROTOC=protoc
PROTO_DIR=api
GEN_DIR=internal/common/genproto
GO_OUT=--go_out=$(GEN_DIR) --go_opt=paths=source_relative
GRPC_OUT=--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative

.PHONY: proto
proto:
	@echo "Generating proto files..."
	@mkdir -p $(GEN_DIR)
	$(PROTOC) $(GO_OUT) $(GRPC_OUT) \
		-I=$(PROTO_DIR) \
		$(PROTO_DIR)/order.proto \
		$(PROTO_DIR)/stock.proto
	@echo "Proto files generated successfully!"

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	rm -rf $(GEN_DIR)
	@echo "Clean complete!"

.PHONY: all
all: proto

