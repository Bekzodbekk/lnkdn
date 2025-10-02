package repository

import (
	"context"
	"like-service/genproto/likepb"
)

type ILikeService interface {
	ToggleLike(ctx context.Context, req *likepb.ToggleLikeReq) (*likepb.ToggleLikeResp, error)
}
