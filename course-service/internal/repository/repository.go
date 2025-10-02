package repository

import (
	"context"
	"course-service/genproto/coursepb"
)

type ICourseService interface {
	CreateCourse(ctx context.Context, req *coursepb.CreateCourseReq) (*coursepb.CreateCourseResp, error)
	UpdateCourse(ctx context.Context, req *coursepb.UpdateCourseReq) (*coursepb.UpdateCourseResp, error)
	DeleteCourse(ctx context.Context, req *coursepb.DeleteCourseReq) (*coursepb.DeleteCourseResp, error)
	GetCourseById(ctx context.Context, req *coursepb.GetCourseByIdReq) (*coursepb.GetCourseByIdResp, error)
	GetCourses(ctx context.Context, req *coursepb.GetCoursesReq) (*coursepb.GetCoursesResp, error)
	GetCourseByUserId(ctx context.Context, req *coursepb.GetCourseByUserIdReq) (*coursepb.GetCourseByUserIdResp, error)

	CreateLesson(ctx context.Context, req *coursepb.CreateLessonReq) (*coursepb.CreateLessonResp, error)
	UpdateLesson(ctx context.Context, req *coursepb.UpdateLessonReq) (*coursepb.UpdateLessonResp, error)
	DeleteLesson(ctx context.Context, req *coursepb.DeleteLessonReq) (*coursepb.DeleteLessonResp, error)
	GetLessonById(ctx context.Context, req *coursepb.GetLessonByIdReq) (*coursepb.GetLessonByIdResp, error)
	GetLessonsByCourseId(ctx context.Context, req *coursepb.GetLessonsByCourseIdReq) (*coursepb.GetLessonsByCourseIdResp, error)
}
