package handlers

import (
	"api-gateway/genproto/authpb"

	"github.com/gin-gonic/gin"
)

func (h *HandlerSt) SignIn(ctx *gin.Context) {
	req := authpb.SignInReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.SignIn(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) ForgetPasswordSendCodeEmail(ctx *gin.Context) {
	req := authpb.ForgetPasswordSendCodeEmailReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.ForgetPasswordSendCodeEmail(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) ForgetPasswordCheckCode(ctx *gin.Context) {
	req := authpb.ForgetPasswordCheckCodeReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.ForgetPasswordCheckCode(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) ForgetPasswordUpdatePassword(ctx *gin.Context) {
	req := authpb.ForgetPasswordUpdatePasswordReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.ForgetPasswordUpdatePassword(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) UpdatePassword(ctx *gin.Context) {
	req := authpb.UpdatePasswordReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.UpdatePassword(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) CheckCodeAfterUpdatePassword(ctx *gin.Context) {
	req := authpb.CheckCodeAfterUpdatePasswordReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.CheckCodeAfterUpdatePassword(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) CreateUser(ctx *gin.Context) {
	req := authpb.CreateUserReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	req := authpb.UpdateUserReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	req.Id = id

	resp, err := h.Service.UpdateUser(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	req := authpb.DeleteUserReq{}
	req.Id = id

	resp, err := h.Service.DeleteUser(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	req := authpb.GetUserByIdReq{}
	req.Id = id

	resp, err := h.Service.GetUserById(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerSt) GetUsers(ctx *gin.Context) {
	req := authpb.GetUsersReq{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.Service.GetUsers(ctx, &req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
