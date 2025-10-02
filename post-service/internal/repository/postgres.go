package repository

import (
	"context"
	"database/sql"
	"post-service/genproto/postpb"
	"post-service/storage"
	"strconv"
)

type PostRepo struct {
	db      *sql.DB
	queries *storage.Queries
}

func NewPostRepo(
	db *sql.DB,
	queries *storage.Queries,
) IPostService {
	return &PostRepo{
		db:      db,
		queries: queries,
	}
}

func (p *PostRepo) CreatePost(ctx context.Context, req *postpb.CreatePostReq) (*postpb.CreatePostResp, error) {
	ID, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}

	err = p.queries.CreatePost(ctx, storage.CreatePostParams{
		UserID:  int32(ID),
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &postpb.CreatePostResp{
		Status:  true,
		Message: "Create Post Successfully",
	}, nil
}
func (p *PostRepo) UpdatePost(ctx context.Context, req *postpb.UpdatePostReq) (*postpb.UpdatePostResp, error) {
	ID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = p.queries.UpdatePost(ctx, storage.UpdatePostParams{
		ID:      int32(ID),
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &postpb.UpdatePostResp{
		Status:  true,
		Message: "post update successfully",
	}, nil
}
func (p *PostRepo) DeletePost(ctx context.Context, req *postpb.DeletePostReq) (*postpb.DeletePostResp, error) {
	ID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = p.queries.DeletePost(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	return &postpb.DeletePostResp{
		Status:  true,
		Message: "Post Delete Successfully",
	}, nil
}
func (p *PostRepo) GetPostById(ctx context.Context, req *postpb.GetPostByIdReq) (*postpb.GetPostByIdResp, error) {
	ID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	resp, err := p.queries.GetPostById(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	return &postpb.GetPostByIdResp{
		Status:  true,
		Message: "Get Post By Id Successfully",
		Post: &postpb.Post{
			Id:      strconv.Itoa(int(resp.ID)),
			UserId:  strconv.Itoa(int(resp.UserID)),
			Content: resp.Content,
		},
	}, nil
}
func (p *PostRepo) GetAllPosts(ctx context.Context, req *postpb.GetAllPostsReq) (*postpb.GetAllPostsResp, error) {
	resp, err := p.queries.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}
	posts := []*postpb.Post{}

	for _, post := range resp {
		pst := postpb.Post{
			Id:      strconv.Itoa(int(post.ID)),
			UserId:  strconv.Itoa(int(post.UserID)),
			Content: post.Content,
		}
		posts = append(posts, &pst)
	}
	return &postpb.GetAllPostsResp{
		Status:  true,
		Message: "Get All posts successfully",
		Posts:   posts,
	}, nil
}
func (p *PostRepo) GetAllPostsByUserId(ctx context.Context, req *postpb.GetAllPostsByUserIdReq) (*postpb.GetAllPostsByUserIdResp, error) {
	ID, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}
	resp, err := p.queries.GetPostByUserId(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	posts := []*postpb.Post{}

	for _, post := range resp {
		pst := postpb.Post{
			Id:      strconv.Itoa(int(post.ID)),
			UserId:  strconv.Itoa(int(post.UserID)),
			Content: post.Content,
		}
		posts = append(posts, &pst)
	}

	return &postpb.GetAllPostsByUserIdResp{
		Status:  true,
		Message: "Get posts by user id successfully",
		Posts:   posts,
	}, nil
}
