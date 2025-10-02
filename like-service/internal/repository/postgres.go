package repository

import (
	"context"
	"like-service/genproto/likepb"
	"like-service/storage"
	"strconv"
)

type LikeRepo struct {
	queries *storage.Queries
}

func NewLikeRepo(queries *storage.Queries) ILikeService {
	return &LikeRepo{
		queries: queries,
	}
}

func (l *LikeRepo) ToggleLike(ctx context.Context, req *likepb.ToggleLikeReq) (*likepb.ToggleLikeResp, error) {
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}

	postId, err := strconv.Atoi(req.PostId)
	if err != nil {
		return nil, err
	}

	checkLike, err := l.queries.CheckLike(ctx, storage.CheckLikeParams{
		UserID: int32(userId),
		PostID: int32(postId),
	})
	if err != nil {
		return nil, err
	}

	if checkLike {
		err = l.queries.RemoveLike(ctx, storage.RemoveLikeParams{
			UserID: int32(userId),
			PostID: int32(postId),
		})
		if err != nil {
			return nil, err
		}
	} else {
		err = l.queries.AddLike(ctx, storage.AddLikeParams{
			UserID: int32(userId),
			PostID: int32(postId),
		})
		if err != nil {
			return nil, err
		}
	}
	return &likepb.ToggleLikeResp{
		Status:  true,
		Message: "Like changed",
	}, nil
}
