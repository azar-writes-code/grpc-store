protoc:
	protoc --go_out=./pb --go-grpc_out=./pb proto/*.proto

clean:
	rm -rf pb/*

server:
	go run main.go