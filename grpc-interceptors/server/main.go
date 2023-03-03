package main

import (
	"context"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"ingens/configs"
	"ingens/service"
	"ingens/tarain"
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("----> Unary Interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("----> Stream Interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

const (
	secretkey     = "secret"
	tokenDuration = 2 * time.Hour
)

var mongoClient *mongo.Client

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin123", "secret", "admin")
	if err != nil {
		return err
	}

	return createUser(userStore, "user123", "secret", "user")
}

func createUser(userStore service.UserStore, username string, password string, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}
func init() {
	mongoClient = configs.ConnectDB()
}
func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	userStore := service.NewInMemoryUserStore(mongoClient)
	jwtManager := service.NewJWTManager(secretkey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)
	err1 := seedUsers(userStore)
	if err1 != nil {
		log.Fatalf("cannot seed users")
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)
	tarain.RegisterAuthServiceServer(s, authServer)
	reflection.Register(s)
	log.Printf("Starting server in port :%d\n", 8000)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
