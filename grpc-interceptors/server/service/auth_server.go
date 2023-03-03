package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"ingens/tarain"
)

type AuthServer struct {
	userStore  UserStore
	jwtManager *JWTManager
	tarain.UnimplementedAuthServiceServer
}

func NewAuthServer(userStore UserStore,
	jwtManager *JWTManager) *AuthServer {
	return &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}

func (server *AuthServer) Login(ctx context.Context, req *tarain.LoginRequest) (*tarain.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find the user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "Incorrect username/password ")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	res := &tarain.LoginResponse{
		AccessToken: token,
	}
	return res, nil
}
