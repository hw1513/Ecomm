#!/bin/bash

# 1. 检查是否在项目根目录
if [[ ! -f "./scripts/generate.sh" ]]; then
  echo "请在项目根目录下运行此脚本！"
  exit 1
fi

# proto 配置：key 为输出子目录，value 为 proto 文件的相对路径
declare -A protos
protos["orderpb"]="./api/orderpb/order.proto"
protos["stockpb"]="./api/stockpb/stock.proto"
PROTO_OUT="./internal/common/genproto"
if [ -d "$PROTO_OUT" ]; then
    rm -rf "$PROTO_OUT"/* 2>/dev/null
    echo "$PROTO_OUT 已清空"
  else
    mkdir -p "$PROTO_OUT"
    echo "$PROTO_OUT 已创建"
  fi


for pb in "${!protos[@]}"; do
  PROTO_SRC="${protos[$pb]}"

  # 取 proto 文件所在目录作为 -I 路径
  PROTO_DIR="./api"

  protoc --go_out="$PROTO_OUT" --go_opt=paths=source_relative \
         --go-grpc_out="$PROTO_OUT" --go-grpc_opt=paths=source_relative \
         -I"$PROTO_DIR" \
         "$PROTO_SRC"

  echo "$(basename "$PROTO_SRC") 已生成到 $PROTO_OUT"
done


OPENAPI_FILE="./api/order.openapi.yaml"
OPENAPI_OUT="./internal/common/genrest"

# 创建目标目录
mkdir -p "$OPENAPI_OUT"

# 生成 types（数据结构）
oapi-codegen -generate types -o "$OPENAPI_OUT/types.gen.go" -package rest "$OPENAPI_FILE"

# 生成 server（接口框架）
oapi-codegen -generate gin -o "$OPENAPI_OUT/server.gen.go" -package rest "$OPENAPI_FILE"

# 生成 client（可选）
# oapi-codegen -generate client -o "$OPENAPI_OUT/client.gen.go" -package rest "$OPENAPI_FILE"

echo "OpenAPI 代码已生成到 $OPENAPI_OUT"