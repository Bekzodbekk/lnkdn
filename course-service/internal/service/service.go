package service

import (
	"context"
	"course-service/genproto/coursepb"
	"course-service/internal/repository"
)

type Service struct {
	coursepb.UnimplementedCourseServiceServer
	repo repository.ICourseService
}

func NewService(repo repository.ICourseService) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateCourse(ctx context.Context, req *coursepb.CreateCourseReq) (*coursepb.CreateCourseResp, error) {
	return s.repo.CreateCourse(ctx, req)
}
func (s *Service) UpdateCourse(ctx context.Context, req *coursepb.UpdateCourseReq) (*coursepb.UpdateCourseResp, error) {
	return s.repo.UpdateCourse(ctx, req)
}
func (s *Service) DeleteCourse(ctx context.Context, req *coursepb.DeleteCourseReq) (*coursepb.DeleteCourseResp, error) {
	return s.repo.DeleteCourse(ctx, req)
}
func (s *Service) GetCourseById(ctx context.Context, req *coursepb.GetCourseByIdReq) (*coursepb.GetCourseByIdResp, error) {
	return s.repo.GetCourseById(ctx, req)
}
func (s *Service) GetCourses(ctx context.Context, req *coursepb.GetCoursesReq) (*coursepb.GetCoursesResp, error) {
	return s.repo.GetCourses(ctx, req)
}
func (s *Service) GetCourseByUserId(ctx context.Context, req *coursepb.GetCourseByUserIdReq) (*coursepb.GetCourseByUserIdResp, error) {
	return s.repo.GetCourseByUserId(ctx, req)
}

func (s *Service) CreateLesson(ctx context.Context, req *coursepb.CreateLessonReq) (*coursepb.CreateLessonResp, error) {
	return s.repo.CreateLesson(ctx, req)
}
func (s *Service) UpdateLesson(ctx context.Context, req *coursepb.UpdateLessonReq) (*coursepb.UpdateLessonResp, error) {
	return s.repo.UpdateLesson(ctx, req)
}
func (s *Service) DeleteLesson(ctx context.Context, req *coursepb.DeleteLessonReq) (*coursepb.DeleteLessonResp, error) {
	return s.repo.DeleteLesson(ctx, req)
}
func (s *Service) GetLessonById(ctx context.Context, req *coursepb.GetLessonByIdReq) (*coursepb.GetLessonByIdResp, error) {
	return s.repo.GetLessonById(ctx, req)
}
func (s *Service) GetLessonsByCourseId(ctx context.Context, req *coursepb.GetLessonsByCourseIdReq) (*coursepb.GetLessonsByCourseIdResp, error) {
	return s.repo.GetLessonsByCourseId(ctx, req)
}
