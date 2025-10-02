package service

import (
	"auth-service/genproto/authpb"
	"auth-service/internal/pkg/config"
	"auth-service/internal/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type RunService struct {
	srv *service.Service
}

func NewRunService(srv *service.Service) *RunService {
	return &RunService{
		srv: srv,
	}
}

func (r *RunService) Run(cfg *config.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.AuthServiceHost, cfg.AuthServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	newServer := grpc.NewServer()
	authpb.RegisterUserServiceServer(newServer, r.srv)
	return newServer.Serve(listener)
}
