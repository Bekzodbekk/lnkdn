package handlers

import (
	"api-gateway/genproto/commentpb"

	"github.com/gin-gonic/gin"
)

func (h *HandlerSt) CreateComment(ctx *gin.Context) {
	req := commentpb.CreateCommentReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.CreateComment(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerSt) UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	req := commentpb.UpdateCommentReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	req.Id = id

	resp, err := h.Service.UpdateComment(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerSt) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.Service.DeleteComment(ctx, &commentpb.DeleteCommentReq{
		Id: id,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerSt) GetCommentByUserId(ctx *gin.Context) {
	user_id := ctx.Param("user_id")

	resp, err := h.Service.GetCommentByUserId(ctx, &commentpb.GetCommentByUserIdReq{
		UserId: user_id,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerSt) GetCommentByPostId(ctx *gin.Context) {
	post_id := ctx.Param("post_id")

	resp, err := h.Service.GetCommentByPostId(ctx, &commentpb.GetCommentByPostIdReq{
		PostId: post_id,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
