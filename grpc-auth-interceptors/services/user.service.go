package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/MrAzharuddin/grpc-auth-interceptors/configs"
	"github.com/MrAzharuddin/grpc-auth-interceptors/models"
	"github.com/MrAzharuddin/grpc-auth-interceptors/pb"
	"github.com/MrAzharuddin/grpc-auth-interceptors/utils"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	client *mongo.Client
}

type UserData struct {
	Id        primitive.ObjectID `bson:"_id, omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func NewAuthServer(client *mongo.Client) *Server {
	return &Server{
		client: client,
	}
}

func (server *Server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	if (req.GetName() == "") || (req.GetEmail() == "") || (req.GetPassword() == "") {
		return nil, fmt.Errorf("Name/Email/Password cannot be empty")
	}

	userCollection := configs.GetCollection(server.client, "users")

	hashPassword := utils.HashPassword(req.GetPassword())
	user := models.User{
		ID:        primitive.NewObjectID(),
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Password:  hashPassword,
		Verified:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.InsertedID)

	return &pb.SignUpResponse{
		User: &pb.User{
			Id:        user.ID.Hex(),
			Name:      req.GetName(),
			Email:     req.GetEmail(),
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}

func (server *Server) Ping(cont context.Context, req *pb.PingRequest) (*pb.PongResponse, error) {
	return &pb.PongResponse{
		Name: "Hello World",
	}, nil
}

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	
	return nil, nil
}
