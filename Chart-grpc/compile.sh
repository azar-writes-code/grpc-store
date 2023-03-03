#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=./server --go-grpc_out=./server proto/*.proto
protoc --go_out=./client --go-grpc_out=./client proto/*.proto
protoc --go_out=./exp --go-grpc_out=./exp proto/*.proto

protoc -I =. */*.proto --js_out=import_style=commonjs,binary:./nextjs/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./nextjs/