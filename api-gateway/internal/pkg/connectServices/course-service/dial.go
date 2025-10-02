package courseservice

import (
	"api-gateway/genproto/coursepb"
	"api-gateway/internal/pkg/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithCourseService(cfg config.Config) (*coursepb.CourseServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.CourseService.Host, cfg.CourseService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	courseServiceClient := coursepb.NewCourseServiceClient(conn)
	return &courseServiceClient, nil
}
