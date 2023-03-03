#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=./server --go-grpc_out=./server proto/*.proto

protoc --go_out=./client --go-grpc_out=./client proto/*.proto

protoc -I =. */*.proto --js_out=import_style=commonjs,binary:./react-app/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./react-app/src/