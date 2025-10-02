package repository

import (
	"comment-service/genproto/commentpb"
	"context"
)

type ICommentService interface {
	CreateComment(ctx context.Context, req *commentpb.CreateCommentReq) (*commentpb.CreateCommentResp, error)
	UpdateComment(ctx context.Context, req *commentpb.UpdateCommentReq) (*commentpb.UpdateCommentResp, error)
	DeleteComment(ctx context.Context, req *commentpb.DeleteCommentReq) (*commentpb.DeleteCommentResp, error)
	GetCommentByUserId(ctx context.Context, req *commentpb.GetCommentByUserIdReq) (*commentpb.GetCommentByUserIdResp, error)
	GetCommentByPostId(ctx context.Context, req *commentpb.GetCommentByPostIdReq) (*commentpb.GetCommentByPostIdResp, error)
}
