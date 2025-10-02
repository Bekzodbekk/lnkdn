package handlers

import (
	"api-gateway/genproto/likepb"

	"github.com/gin-gonic/gin"
)

func (h *HandlerSt) ToggleLike(ctx *gin.Context) {
	req := likepb.ToggleLikeReq{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.ToggleLike(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
