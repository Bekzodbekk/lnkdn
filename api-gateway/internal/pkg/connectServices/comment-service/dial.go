package commentservice

import (
	"api-gateway/genproto/commentpb"
	"api-gateway/internal/pkg/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithCommentService(cfg config.Config) (*commentpb.CommentServiceClient, error) {
	target := fmt.Sprintf("%s:%d", cfg.CommentService.Host, cfg.CommentService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := commentpb.NewCommentServiceClient(conn)
	return &client, nil
}
