package main

import (
	"log"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
	pb "grpc/proto"
	ch "grpc/clientHelpers"
)

const (
	port = ":8080"
)

func main()  {
	conn, err := grpc.Dial("localhost" + port, grpc.WithInsecure()) //grpc.WithTransportCredentials(insecure.NewCredentials())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	ch.CallSayHello(client)
	ch.CallSayHelloServerStream(client, names)
	ch.CallSayHelloClientStream(client, names)
	ch.CallSayHelloBidirectionalStream(client, names)
}