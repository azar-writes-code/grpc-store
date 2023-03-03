package clientHelpers

import (
	"context"
	"log"
	"time"
	pb "grpc/proto"
)

func CallSayHello(client pb.GreetServiceClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}

