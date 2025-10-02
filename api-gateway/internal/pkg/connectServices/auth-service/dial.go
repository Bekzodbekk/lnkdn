package authservice

import (
	"api-gateway/genproto/authpb"
	"api-gateway/internal/pkg/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithAuthService(cfg *config.Config) (*authpb.UserServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.AuthService.Host, cfg.AuthService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := authpb.NewUserServiceClient(conn)
	return &userServiceClient, nil
}
