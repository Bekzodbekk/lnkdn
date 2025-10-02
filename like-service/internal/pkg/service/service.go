package service

import (
	"fmt"
	"like-service/genproto/likepb"
	"like-service/internal/pkg/config"
	"like-service/internal/service"
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
func (r *RunService) RUN(cfg *config.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.LikeServiceHost, cfg.LikeServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return nil
	}

	server := grpc.NewServer()
	likepb.RegisterLikeServiceServer(server, r.srv)

	return server.Serve(listener)
}
