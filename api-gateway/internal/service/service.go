package service

import (
	"api-gateway/genproto/authpb"
	"api-gateway/genproto/commentpb"
	"api-gateway/genproto/coursepb"
	"api-gateway/genproto/likepb"
	"api-gateway/genproto/postpb"
	"context"
)

type ServiceRepositoryClient struct {
	authClient    authpb.UserServiceClient
	postClient    postpb.PostServiceClient
	likeClient    likepb.LikeServiceClient
	commentClient commentpb.CommentServiceClient
	courseClient  coursepb.CourseServiceClient
}

func NewServiceRepositoryClient(
	authClient authpb.UserServiceClient,
	postClient postpb.PostServiceClient,
	likeClient likepb.LikeServiceClient,
	commentClient commentpb.CommentServiceClient,
	courseClient coursepb.CourseServiceClient,

) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		authClient:    authClient,
		postClient:    postClient,
		likeClient:    likeClient,
		commentClient: commentClient,
		courseClient:  courseClient,
	}
}

// ! AUTH SERVICE METHODS
func (s *ServiceRepositoryClient) SignIn(ctx context.Context, req *authpb.SignInReq) (*authpb.SignInResp, error) {
	return s.authClient.SignIn(ctx, req)
}
func (s *ServiceRepositoryClient) ForgetPasswordSendCodeEmail(ctx context.Context, req *authpb.ForgetPasswordSendCodeEmailReq) (*authpb.ForgetPasswordSendCodeEmailResp, error) {
	return s.authClient.ForgetPasswordSendCodeEmail(ctx, req)
}
func (s *ServiceRepositoryClient) ForgetPasswordCheckCode(ctx context.Context, req *authpb.ForgetPasswordCheckCodeReq) (*authpb.ForgetPasswordCheckCodeResp, error) {
	return s.authClient.ForgetPasswordCheckCode(ctx, req)
}
func (s *ServiceRepositoryClient) ForgetPasswordUpdatePassword(ctx context.Context, req *authpb.ForgetPasswordUpdatePasswordReq) (*authpb.ForgetPasswordUpdatePasswordResp, error) {
	return s.authClient.ForgetPasswordUpdatePassword(ctx, req)
}
func (s *ServiceRepositoryClient) UpdatePassword(ctx context.Context, req *authpb.UpdatePasswordReq) (*authpb.UpdateUserResp, error) {
	return s.authClient.UpdatePassword(ctx, req)
}
func (s *ServiceRepositoryClient) CheckCodeAfterUpdatePassword(ctx context.Context, req *authpb.CheckCodeAfterUpdatePasswordReq) (*authpb.CheckCodeAfterUpdatePasswordResp, error) {
	return s.authClient.CheckCodeAfterUpdatePassword(ctx, req)
}

func (s *ServiceRepositoryClient) CreateUser(ctx context.Context, req *authpb.CreateUserReq) (*authpb.CreateUserResp, error) {
	return s.authClient.CreateUser(ctx, req)
}
func (s *ServiceRepositoryClient) UpdateUser(ctx context.Context, req *authpb.UpdateUserReq) (*authpb.UpdateUserResp, error) {
	return s.authClient.UpdateUser(ctx, req)
}
func (s *ServiceRepositoryClient) DeleteUser(ctx context.Context, req *authpb.DeleteUserReq) (*authpb.DeleteUserResp, error) {
	return s.authClient.DeleteUser(ctx, req)
}
func (s *ServiceRepositoryClient) GetUserById(ctx context.Context, req *authpb.GetUserByIdReq) (*authpb.GetUserByIdResp, error) {
	return s.authClient.GetUserById(ctx, req)
}
func (s *ServiceRepositoryClient) GetUsers(ctx context.Context, req *authpb.GetUsersReq) (*authpb.GetUsersResp, error) {
	return s.authClient.GetUsers(ctx, req)
}

// ! POST SERVICE METHODS
func (s *ServiceRepositoryClient) CreatePost(ctx context.Context, req *postpb.CreatePostReq) (*postpb.CreatePostResp, error) {
	return s.postClient.CreatePost(ctx, req)
}
func (s *ServiceRepositoryClient) UpdatePost(ctx context.Context, req *postpb.UpdatePostReq) (*postpb.UpdatePostResp, error) {
	return s.postClient.UpdatePost(ctx, req)
}
func (s *ServiceRepositoryClient) DeletePost(ctx context.Context, req *postpb.DeletePostReq) (*postpb.DeletePostResp, error) {
	return s.postClient.DeletePost(ctx, req)
}
func (s *ServiceRepositoryClient) GetPostById(ctx context.Context, req *postpb.GetPostByIdReq) (*postpb.GetPostByIdResp, error) {
	return s.postClient.GetPostById(ctx, req)
}
func (s *ServiceRepositoryClient) GetAllPosts(ctx context.Context, req *postpb.GetAllPostsReq) (*postpb.GetAllPostsResp, error) {
	return s.postClient.GetAllPosts(ctx, req)
}
func (s *ServiceRepositoryClient) GetAllPostsByUserId(ctx context.Context, req *postpb.GetAllPostsByUserIdReq) (*postpb.GetAllPostsByUserIdResp, error) {
	return s.postClient.GetAllPostsByUserId(ctx, req)
}

// ! LIKE SERVICE METHODS
func (s *ServiceRepositoryClient) ToggleLike(ctx context.Context, req *likepb.ToggleLikeReq) (*likepb.ToggleLikeResp, error) {
	return s.likeClient.ToggleLike(ctx, req)
}

// ! COMMENT SERVICE METHODS
func (s *ServiceRepositoryClient) CreateComment(ctx context.Context, req *commentpb.CreateCommentReq) (*commentpb.CreateCommentResp, error) {
	return s.commentClient.CreateComment(ctx, req)
}
func (s *ServiceRepositoryClient) UpdateComment(ctx context.Context, req *commentpb.UpdateCommentReq) (*commentpb.UpdateCommentResp, error) {
	return s.commentClient.UpdateComment(ctx, req)
}
func (s *ServiceRepositoryClient) DeleteComment(ctx context.Context, req *commentpb.DeleteCommentReq) (*commentpb.DeleteCommentResp, error) {
	return s.commentClient.DeleteComment(ctx, req)
}
func (s *ServiceRepositoryClient) GetCommentByUserId(ctx context.Context, req *commentpb.GetCommentByUserIdReq) (*commentpb.GetCommentByUserIdResp, error) {
	return s.commentClient.GetCommentByUserId(ctx, req)
}
func (s *ServiceRepositoryClient) GetCommentByPostId(ctx context.Context, req *commentpb.GetCommentByPostIdReq) (*commentpb.GetCommentByPostIdResp, error) {
	return s.commentClient.GetCommentByPostId(ctx, req)
}

// ! COURSE SERVICE METHODS
func (s *ServiceRepositoryClient) CreateCourse(ctx context.Context, req *coursepb.CreateCourseReq) (*coursepb.CreateCourseResp, error){
	return s.courseClient.CreateCourse(ctx, req)
}
func (s *ServiceRepositoryClient) UpdateCourse(ctx context.Context, req *coursepb.UpdateCourseReq) (*coursepb.UpdateCourseResp, error){
	return s.courseClient.UpdateCourse(ctx, req)
}
func (s *ServiceRepositoryClient) DeleteCourse(ctx context.Context, req *coursepb.DeleteCourseReq) (*coursepb.DeleteCourseResp, error){
	return s.courseClient.DeleteCourse(ctx, req)
}
func (s *ServiceRepositoryClient) GetCourseById(ctx context.Context, req *coursepb.GetCourseByIdReq) (*coursepb.GetCourseByIdResp, error){
	return s.courseClient.GetCourseById(ctx, req)
}
func (s *ServiceRepositoryClient) GetCourses(ctx context.Context, req *coursepb.GetCoursesReq) (*coursepb.GetCoursesResp, error){
	return s.courseClient.GetCourses(ctx, req)
}
func (s *ServiceRepositoryClient) GetCourseByUserId(ctx context.Context, req *coursepb.GetCourseByUserIdReq) (*coursepb.GetCourseByUserIdResp, error){
	return s.courseClient.GetCourseByUserId(ctx, req)
}

func (s *ServiceRepositoryClient) CreateLesson(ctx context.Context, req *coursepb.CreateLessonReq) (*coursepb.CreateLessonResp, error){
	return s.courseClient.CreateLesson(ctx, req)
}
func (s *ServiceRepositoryClient) UpdateLesson(ctx context.Context, req *coursepb.UpdateLessonReq) (*coursepb.UpdateLessonResp, error){
	return s.courseClient.UpdateLesson(ctx, req)
}
func (s *ServiceRepositoryClient) DeleteLesson(ctx context.Context, req *coursepb.DeleteLessonReq) (*coursepb.DeleteLessonResp, error){
	return s.courseClient.DeleteLesson(ctx, req)
}
func (s *ServiceRepositoryClient) GetLessonById(ctx context.Context, req *coursepb.GetLessonByIdReq) (*coursepb.GetLessonByIdResp, error){
	return s.courseClient.GetLessonById(ctx, req)
}
func (s *ServiceRepositoryClient) GetLessonsByCourseId(ctx context.Context, req *coursepb.GetLessonsByCourseIdReq) (*coursepb.GetLessonsByCourseIdResp, error){
	return s.courseClient.GetLessonsByCourseId(ctx, req)
}
