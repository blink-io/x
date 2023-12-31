#!/bin/bash


TRANS_PATH=github.com/blink-io/x/kratos/v2/transport/http
EXTERN_TMPL=/data/projects/open/x/kratos/v2/cmd/protoc-gen-go-http/httpTemplate-Blink.tpl
THIRD_PARTY_PROTOS=/data/projects/open/x/third_party/

protoc --proto_path=. \
  --proto_path=$THIRD_PARTY_PROTOS \
  --go-http_out=extern_template=${EXTERN_TMPL},transport_path=${TRANS_PATH}:. \
  --plugin /home/heisonyee/go/bin/protoc-gen-go-http \
  ./metadata.proto