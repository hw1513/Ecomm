#!/bin/bash

# 确保脚本在错误时退出
set -e

# 检查必要的工具是否安装
command -v protoc >/dev/null 2>&1 || { echo "需要安装 protoc"; exit 1; }
command -v go >/dev/null 2>&1 || { echo "需要安装 go"; exit 1; }

# 安装必要的 Go 工具
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 创建必要的目录
mkdir -p internal/order/api
mkdir -p internal/order/rest

# 生成 gRPC 代码
echo "Generating gRPC code..."
protoc -I . \
  --go_out internal/order/api --go_opt paths=source_relative \
  --go-grpc_out internal/order/api --go-grpc_opt paths=source_relative \
  api/order.proto

# 生成 REST API 代码
echo "Generating REST API code..."
oapi-codegen -package rest api/order.openapi.yaml > internal/order/rest/rest.go

echo "Code generation completed successfully!" 