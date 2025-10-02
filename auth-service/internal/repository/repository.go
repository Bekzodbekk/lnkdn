package repository

import (
	"auth-service/genproto/authpb"
	"context"
)

type IUserRepository interface {
	SignIn(ctx context.Context, req *authpb.SignInReq) (*authpb.SignInResp, error)
	ForgetPasswordSendCodeEmail(ctx context.Context, req *authpb.ForgetPasswordSendCodeEmailReq) (*authpb.ForgetPasswordSendCodeEmailResp, error)
	ForgetPasswordCheckCode(ctx context.Context, req *authpb.ForgetPasswordCheckCodeReq) (*authpb.ForgetPasswordCheckCodeResp, error)
	ForgetPasswordUpdatePassword(ctx context.Context, req *authpb.ForgetPasswordUpdatePasswordReq) (*authpb.ForgetPasswordUpdatePasswordResp, error)
	UpdatePassword(ctx context.Context, req *authpb.UpdatePasswordReq) (*authpb.UpdateUserResp, error)
	CheckCodeAfterUpdatePassword(ctx context.Context, req *authpb.CheckCodeAfterUpdatePasswordReq) (*authpb.CheckCodeAfterUpdatePasswordResp, error)

	CreateUser(ctx context.Context, req *authpb.CreateUserReq) (*authpb.CreateUserResp, error)
	UpdateUser(ctx context.Context, req *authpb.UpdateUserReq) (*authpb.UpdateUserResp, error)
	DeleteUser(ctx context.Context, req *authpb.DeleteUserReq) (*authpb.DeleteUserResp, error)
	GetUserById(ctx context.Context, req *authpb.GetUserByIdReq) (*authpb.GetUserByIdResp, error)
	GetUsers(ctx context.Context, req *authpb.GetUsersReq) (*authpb.GetUsersResp, error)
}
