package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/MrAzharuddin/grpc-auth-interceptors/configs"
	"github.com/MrAzharuddin/grpc-auth-interceptors/pb"
	"github.com/MrAzharuddin/grpc-auth-interceptors/services"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error listening: %e", err)
		panic(err)
	}
	db := configs.ConnectDB()
	authserver := services.NewAuthServer(db)
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, authserver)
	reflection.Register(s)

	log.Printf("Starting server in port :%d\n", 9000)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
