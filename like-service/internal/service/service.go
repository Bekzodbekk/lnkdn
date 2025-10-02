package service

import (
	"context"
	"like-service/genproto/likepb"
	"like-service/internal/repository"
)

type Service struct {
	likepb.UnimplementedLikeServiceServer
	repo repository.ILikeService
}

func NewService(repo repository.ILikeService) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) ToggleLike(ctx context.Context, req *likepb.ToggleLikeReq) (*likepb.ToggleLikeResp, error) {
	return s.repo.ToggleLike(ctx, req)
}
