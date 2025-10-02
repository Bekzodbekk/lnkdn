package repository

import (
	"context"
	"post-service/genproto/postpb"
)

type IPostService interface {
	CreatePost(ctx context.Context, req *postpb.CreatePostReq) (*postpb.CreatePostResp, error)
	UpdatePost(ctx context.Context, req *postpb.UpdatePostReq) (*postpb.UpdatePostResp, error)
	DeletePost(ctx context.Context, req *postpb.DeletePostReq) (*postpb.DeletePostResp, error)
	GetPostById(ctx context.Context, req *postpb.GetPostByIdReq) (*postpb.GetPostByIdResp, error)
	GetAllPosts(ctx context.Context, req *postpb.GetAllPostsReq) (*postpb.GetAllPostsResp, error)
	GetAllPostsByUserId(ctx context.Context, req *postpb.GetAllPostsByUserIdReq) (*postpb.GetAllPostsByUserIdResp, error)
}
