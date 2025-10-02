package repository

import (
	"comment-service/genproto/commentpb"
	"comment-service/storage"
	"context"
	"database/sql"
	"strconv"
)

type CommentRepo struct {
	db      *sql.DB
	queries *storage.Queries
}

func NewCommentRepo(
	db *sql.DB,
	queries *storage.Queries,
) ICommentService {
	return &CommentRepo{
		db:      db,
		queries: queries,
	}
}

func (c *CommentRepo) CreateComment(ctx context.Context, req *commentpb.CreateCommentReq) (*commentpb.CreateCommentResp, error) {
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}

	postId, err := strconv.Atoi(req.PostId)
	if err != nil {
		return nil, err
	}

	err = c.queries.CreateComment(ctx, storage.CreateCommentParams{
		UserID:  int64(userId),
		PostID:  int64(postId),
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &commentpb.CreateCommentResp{
		Status:  true,
		Message: "Create Comment Successfully",
	}, nil
}

func (c *CommentRepo) UpdateComment(ctx context.Context, req *commentpb.UpdateCommentReq) (*commentpb.UpdateCommentResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = c.queries.UpdateComment(ctx, storage.UpdateCommentParams{
		ID:      int32(id),
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &commentpb.UpdateCommentResp{
		Status:  true,
		Message: "Update Comment successfully",
	}, nil
}
func (c *CommentRepo) DeleteComment(ctx context.Context, req *commentpb.DeleteCommentReq) (*commentpb.DeleteCommentResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = c.queries.DeleteComment(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &commentpb.DeleteCommentResp{
		Status:  true,
		Message: "Delete Comment Successfully",
	}, nil
}
func (c *CommentRepo) GetCommentByUserId(ctx context.Context, req *commentpb.GetCommentByUserIdReq) (*commentpb.GetCommentByUserIdResp, error) {
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}

	resp, err := c.queries.GetCommentByUserId(ctx, int64(userId))
	if err != nil {
		return nil, err
	}

	comments := []*commentpb.Comment{}

	for _, comment := range resp {
		cmnt := commentpb.Comment{
			Id:      strconv.Itoa(int(comment.ID)),
			UserId:  strconv.Itoa(int(comment.UserID)),
			PostId:  strconv.Itoa(int(comment.PostID)),
			Content: comment.Content,
		}
		comments = append(comments, &cmnt)
	}

	return &commentpb.GetCommentByUserIdResp{
		Status:   true,
		Message:  "Get Comment By User Id Successfully",
		Comments: comments,
	}, nil
}
func (c *CommentRepo) GetCommentByPostId(ctx context.Context, req *commentpb.GetCommentByPostIdReq) (*commentpb.GetCommentByPostIdResp, error) {
	postId, err := strconv.Atoi(req.PostId)
	if err != nil {
		return nil, err
	}

	resp, err := c.queries.GetCommentByPostId(ctx, int64(postId))
	if err != nil {
		return nil, err
	}

	comments := []*commentpb.Comment{}

	for _, comment := range resp {
		cmnt := commentpb.Comment{
			Id:      strconv.Itoa(int(comment.ID)),
			UserId:  strconv.Itoa(int(comment.UserID)),
			PostId:  strconv.Itoa(int(comment.PostID)),
			Content: comment.Content,
		}
		comments = append(comments, &cmnt)
	}

	return &commentpb.GetCommentByPostIdResp{
		Status:   true,
		Message:  "Get Comment By Post Id Successfully",
		Comments: comments,
	}, nil
}
