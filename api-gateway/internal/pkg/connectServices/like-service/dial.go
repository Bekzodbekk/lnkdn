package likeservice

import (
	"api-gateway/genproto/likepb"
	"api-gateway/internal/pkg/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithLikeService(cfg config.Config) (*likepb.LikeServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.LikeService.Host, cfg.LikeService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := likepb.NewLikeServiceClient(conn)
	return &client, nil
}
