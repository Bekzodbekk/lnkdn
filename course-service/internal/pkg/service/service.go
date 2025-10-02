package service

import (
	"course-service/genproto/coursepb"
	"course-service/internal/pkg/config"
	"course-service/internal/service"
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

func (r *RunService) RUN(cfg config.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.CourseServiceHost, cfg.CourseServicePort)
	listenn, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	coursepb.RegisterCourseServiceServer(server, r.srv)

	return server.Serve(listenn)
}
