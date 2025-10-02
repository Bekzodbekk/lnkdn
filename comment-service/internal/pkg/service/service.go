package service

import (
	"comment-service/genproto/commentpb"
	"comment-service/internal/pkg/config"
	"comment-service/internal/service"
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
func (r *RunService) RUN(cfg *config.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return nil
	}

	server := grpc.NewServer()
	commentpb.RegisterCommentServiceServer(server, r.srv)

	return server.Serve(listener)
}
