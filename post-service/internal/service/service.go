package service

import (
	"context"
	"post-service/genproto/postpb"
	"post-service/internal/repository"
)

type Service struct {
	postpb.UnimplementedPostServiceServer
	repo repository.IPostService
}

func NewService(repo repository.IPostService) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreatePost(ctx context.Context, req *postpb.CreatePostReq) (*postpb.CreatePostResp, error) {
	return s.repo.CreatePost(ctx, req)
}
func (s *Service) UpdatePost(ctx context.Context, req *postpb.UpdatePostReq) (*postpb.UpdatePostResp, error) {
	return s.repo.UpdatePost(ctx, req)
}
func (s *Service) DeletePost(ctx context.Context, req *postpb.DeletePostReq) (*postpb.DeletePostResp, error) {
	return s.repo.DeletePost(ctx, req)
}
func (s *Service) GetPostById(ctx context.Context, req *postpb.GetPostByIdReq) (*postpb.GetPostByIdResp, error) {
	return s.repo.GetPostById(ctx, req)
}
func (s *Service) GetAllPosts(ctx context.Context, req *postpb.GetAllPostsReq) (*postpb.GetAllPostsResp, error) {
	return s.repo.GetAllPosts(ctx, req)
}
func (s *Service) GetAllPostsByUserId(ctx context.Context, req *postpb.GetAllPostsByUserIdReq) (*postpb.GetAllPostsByUserIdResp, error) {
	return s.repo.GetAllPostsByUserId(ctx, req)
}
