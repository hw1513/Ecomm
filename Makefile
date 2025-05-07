.PHONY: generate build test clean

# 默认目标
all: generate build

# 生成代码
generate:
	@echo "Generating code..."
	@chmod +x scripts/generate.sh
	@./scripts/generate.sh

# 构建项目
build:
	@echo "Building project..."
	@go build -o bin/server cmd/server/main.go

# 运行测试
test:
	@echo "Running tests..."
	@go test -v ./...

# 清理生成的文件
clean:
	@echo "Cleaning generated files..."
	@rm -f internal/order/*.pb.go
	@rm -f internal/order/*.grpc.pb.go
	@rm -f internal/order/*.gw.go
	@rm -f internal/order/rest.go

# 安装依赖
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# 帮助信息
help:
	@echo "Available targets:"
	@echo "  generate  - Generate code from proto and OpenAPI definitions"
	@echo "  build     - Build the project"
	@echo "  test      - Run tests"
	@echo "  clean     - Clean generated files"
	@echo "  deps      - Install dependencies"
	@echo "  help      - Show this help message"