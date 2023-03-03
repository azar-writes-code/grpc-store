gen:
	protoc --go_out=./pb --go_opt=paths=import --go-grpc_out=./pb --go-grpc_opt=paths=import proto/*.proto

path_fix:
	export PATH="$PATH:$(go env GOPATH)/bin"
