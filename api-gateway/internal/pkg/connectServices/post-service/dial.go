package postservice

import (
	"api-gateway/genproto/postpb"
	"api-gateway/internal/pkg/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithPostService(cfg *config.Config) (*postpb.PostServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.PostService.Host, cfg.PostService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	postServiceClient := postpb.NewPostServiceClient(conn)
	return &postServiceClient, nil
}
