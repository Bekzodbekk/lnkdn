package service

import (
	"fmt"
	"net"
	"post-service/genproto/postpb"
	"post-service/internal/pkg/config"
	"post-service/internal/service"

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

func (r *RunService) RUN(cfg *config.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	postpb.RegisterPostServiceServer(server, r.srv)

	return server.Serve(listener)
}
