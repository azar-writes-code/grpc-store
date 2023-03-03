package main

import (
	"log"
	"io"
	pb "grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages[] string

	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil{
			return err
		}

		log.Printf("Got request with name : %v", req.Name)
		messages = append(messages, "Hellooooo "+req.Name)

	}
}