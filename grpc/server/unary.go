package main

import (
	"context"
	pb "grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message:"Hi Sridhar!!",
	}, nil
}