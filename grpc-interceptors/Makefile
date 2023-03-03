install:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

clean:
	rm -rf pb/*.go
	rm -rf server/tarain/*.go

path:
	export PATH="$PATH:$(go env GOPATH)/bin"

protoc:
	protoc --go_out=./pb --go-grpc_out=./pb proto/*.proto
	protoc --go_out=./server/tarain --go-grpc_out=./server/tarain proto/*.proto

protoc_server:
	protoc --go_out=./server/pb --go-grpc_out=./server/pb proto/*.proto

server_p:
	go run server/main.go
