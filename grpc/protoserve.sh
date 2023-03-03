#!/bin/bash

file="proto/*.pb.go"

if [ -f "$file" ] ; then
    rm "$file"
fi
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go-grpc_out=. proto/*.proto