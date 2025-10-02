package service

import (
	"auth-service/genproto/authpb"
	"auth-service/internal/repository"
	"context"
)

type Service struct {
	authpb.UnimplementedUserServiceServer
	repo repository.IUserRepository
}

func NewService(r repository.IUserRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) SignIn(ctx context.Context, req *authpb.SignInReq) (*authpb.SignInResp, error) {
	return s.repo.SignIn(ctx, req)
}
func (s *Service) ForgetPasswordSendCodeEmail(ctx context.Context, req *authpb.ForgetPasswordSendCodeEmailReq) (*authpb.ForgetPasswordSendCodeEmailResp, error) {
	return s.repo.ForgetPasswordSendCodeEmail(ctx, req)
}
func (s *Service) ForgetPasswordCheckCode(ctx context.Context, req *authpb.ForgetPasswordCheckCodeReq) (*authpb.ForgetPasswordCheckCodeResp, error) {
	return s.repo.ForgetPasswordCheckCode(ctx, req)
}
func (s *Service) ForgetPasswordUpdatePassword(ctx context.Context, req *authpb.ForgetPasswordUpdatePasswordReq) (*authpb.ForgetPasswordUpdatePasswordResp, error) {
	return s.repo.ForgetPasswordUpdatePassword(ctx, req)
}
func (s *Service) CreateUser(ctx context.Context, req *authpb.CreateUserReq) (*authpb.CreateUserResp, error) {
	return s.repo.CreateUser(ctx, req)
}
func (s *Service) UpdateUser(ctx context.Context, req *authpb.UpdateUserReq) (*authpb.UpdateUserResp, error) {
	return s.repo.UpdateUser(ctx, req)
}
func (s *Service) DeleteUser(ctx context.Context, req *authpb.DeleteUserReq) (*authpb.DeleteUserResp, error) {
	return s.repo.DeleteUser(ctx, req)
}
func (s *Service) GetUserById(ctx context.Context, req *authpb.GetUserByIdReq) (*authpb.GetUserByIdResp, error) {
	return s.repo.GetUserById(ctx, req)
}
func (s *Service) GetUsers(ctx context.Context, req *authpb.GetUsersReq) (*authpb.GetUsersResp, error) {
	return s.repo.GetUsers(ctx, req)
}
func (s *Service) UpdatePassword(ctx context.Context, req *authpb.UpdatePasswordReq) (*authpb.UpdateUserResp, error) {
	return s.repo.UpdatePassword(ctx, req)
}
func (s *Service) CheckCodeAfterUpdatePassword(ctx context.Context, req *authpb.CheckCodeAfterUpdatePasswordReq) (*authpb.CheckCodeAfterUpdatePasswordResp, error) {
	return s.repo.CheckCodeAfterUpdatePassword(ctx, req)
}
