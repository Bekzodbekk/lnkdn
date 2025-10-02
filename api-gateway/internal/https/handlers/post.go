package handlers

import (
	"api-gateway/genproto/postpb"

	"github.com/gin-gonic/gin"
)

func (h *HandlerSt) CreatePost(ctx *gin.Context) {
	req := postpb.CreatePostReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	resp, err := h.Service.CreatePost(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerSt) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	req := postpb.UpdatePostReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	req.Id = id
	resp, err := h.Service.UpdatePost(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerSt) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := h.Service.DeletePost(ctx, &postpb.DeletePostReq{
		Id: id,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerSt) GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := h.Service.GetPostById(ctx, &postpb.GetPostByIdReq{
		Id: id,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
func (h *HandlerSt) GetAllPosts(ctx *gin.Context) {
	resp, err := h.Service.GetAllPosts(ctx, &postpb.GetAllPostsReq{})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)

}
func (h *HandlerSt) GetAllPostsByUserId(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	resp, err := h.Service.GetAllPostsByUserId(ctx, &postpb.GetAllPostsByUserIdReq{
		UserId: user_id,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
