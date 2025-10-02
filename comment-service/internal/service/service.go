package service

import (
	"comment-service/genproto/commentpb"
	"comment-service/internal/repository"
	"context"
)

type Service struct {
	commentpb.UnimplementedCommentServiceServer
	repo repository.ICommentService
}

func NewService(repo repository.ICommentService) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateComment(ctx context.Context, req *commentpb.CreateCommentReq) (*commentpb.CreateCommentResp, error) {
	return s.repo.CreateComment(ctx, req)
}
func (s *Service) UpdateComment(ctx context.Context, req *commentpb.UpdateCommentReq) (*commentpb.UpdateCommentResp, error) {
	return s.repo.UpdateComment(ctx, req)
}
func (s *Service) DeleteComment(ctx context.Context, req *commentpb.DeleteCommentReq) (*commentpb.DeleteCommentResp, error) {
	return s.repo.DeleteComment(ctx, req)
}
func (s *Service) GetCommentByUserId(ctx context.Context, req *commentpb.GetCommentByUserIdReq) (*commentpb.GetCommentByUserIdResp, error) {
	return s.repo.GetCommentByUserId(ctx, req)
}
func (s *Service) GetCommentByPostId(ctx context.Context, req *commentpb.GetCommentByPostIdReq) (*commentpb.GetCommentByPostIdResp, error) {
	return s.repo.GetCommentByPostId(ctx, req)
}
