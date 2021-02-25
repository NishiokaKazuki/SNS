#!/usr/bin/env bash

set -eu

export PATH="$PATH:$(yarn bin)"

# protoファイルを含むディレクトリの相対パス
PROTO_SRC=../protobuf
# 生成したjs、tsファイルを格納したいディレクトリの相対パス
PROTO_DEST=./src/proto

mkdir -p ${PROTO_DEST}

# protoc-gen-tsのパス
PROTOC_GEN_TS_PATH="$(yarn bin)/protoc-gen-ts"
echo $PROTOC_GEN_TS_PATH
protoc \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs,binary:${PROTO_DEST}" \
    --ts_out="service=grpc-web:${PROTO_DEST}" \
    -I ${PROTO_SRC} $(find ${PROTO_SRC} -name "*.proto")