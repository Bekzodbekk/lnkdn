package repository

import (
	"context"
	"course-service/genproto/coursepb"
	"course-service/storage"
	"database/sql"
	"strconv"
)

type CourseRepo struct {
	db      *sql.DB
	queries *storage.Queries
}

func NewCourseRepo(
	db *sql.DB,
	queries *storage.Queries,
) ICourseService {
	return &CourseRepo{
		db:      db,
		queries: queries,
	}
}

func (c *CourseRepo) CreateCourse(ctx context.Context, req *coursepb.CreateCourseReq) (*coursepb.CreateCourseResp, error) {
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}

	err = c.queries.CreateCourse(ctx, storage.CreateCourseParams{
		UserID:       int32(userId),
		Title:        req.Title,
		Description:  req.Description,
		Category:     req.Category,
		Level:        req.Level,
		ThumbnailUrl: req.ThumbnailUrl,
	})
	if err != nil {
		return nil, err
	}

	return &coursepb.CreateCourseResp{
		Status:  true,
		Message: "Create Course Successfully",
	}, nil
}
func (c *CourseRepo) UpdateCourse(ctx context.Context, req *coursepb.UpdateCourseReq) (*coursepb.UpdateCourseResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = c.queries.UpdateCourse(ctx, storage.UpdateCourseParams{
		ID:           int32(id),
		Title:        req.Title,
		Description:  req.Description,
		Category:     req.Category,
		Level:        req.Level,
		ThumbnailUrl: req.ThumbnailUrl,
	})
	if err != nil {
		return nil, err
	}

	return &coursepb.UpdateCourseResp{
		Status:  true,
		Message: "Course update successfully",
	}, nil
}
func (c *CourseRepo) DeleteCourse(ctx context.Context, req *coursepb.DeleteCourseReq) (*coursepb.DeleteCourseResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = c.queries.DeleteCourse(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &coursepb.DeleteCourseResp{
		Status:  true,
		Message: "Delete Course Successfully",
	}, nil
}
func (c *CourseRepo) GetCourseById(ctx context.Context, req *coursepb.GetCourseByIdReq) (*coursepb.GetCourseByIdResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := c.queries.GetCourseById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &coursepb.GetCourseByIdResp{
		Status:  true,
		Message: "Get Course By Id successfully",
		Course: &coursepb.Course{
			Id:           req.Id,
			UserId:       strconv.Itoa(int(resp.UserID)),
			Title:        resp.Title,
			Description:  resp.Description,
			Category:     resp.Category,
			Level:        resp.Level,
			ThumbnailUrl: resp.ThumbnailUrl,
		},
	}, nil
}
func (c *CourseRepo) GetCourses(ctx context.Context, req *coursepb.GetCoursesReq) (*coursepb.GetCoursesResp, error) {
	resp, err := c.queries.GetCourses(ctx)
	if err != nil {
		return nil, err
	}

	courses := []*coursepb.Course{}

	for _, course := range resp {
		crs := coursepb.Course{
			Id:           strconv.Itoa(int(course.ID)),
			UserId:       strconv.Itoa(int(course.UserID)),
			Title:        course.Title,
			Description:  course.Description,
			Category:     course.Category,
			Level:        course.Level,
			ThumbnailUrl: course.ThumbnailUrl,
		}
		courses = append(courses, &crs)
	}

	return &coursepb.GetCoursesResp{
		Status:  true,
		Message: "Get Courses Successfully",
		Courses: courses,
	}, nil
}
func (c *CourseRepo) GetCourseByUserId(ctx context.Context, req *coursepb.GetCourseByUserIdReq) (*coursepb.GetCourseByUserIdResp, error) {
	
	return nil, nil
}

func (c *CourseRepo) CreateLesson(ctx context.Context, req *coursepb.CreateLessonReq) (*coursepb.CreateLessonResp, error) {
	return nil, nil
}
func (c *CourseRepo) UpdateLesson(ctx context.Context, req *coursepb.UpdateLessonReq) (*coursepb.UpdateLessonResp, error) {
	return nil, nil
}
func (c *CourseRepo) DeleteLesson(ctx context.Context, req *coursepb.DeleteLessonReq) (*coursepb.DeleteLessonResp, error) {
	return nil, nil
}
func (c *CourseRepo) GetLessonById(ctx context.Context, req *coursepb.GetLessonByIdReq) (*coursepb.GetLessonByIdResp, error) {
	return nil, nil
}
func (c *CourseRepo) GetLessonsByCourseId(ctx context.Context, req *coursepb.GetLessonsByCourseIdReq) (*coursepb.GetLessonsByCourseIdResp, error) {
	return nil, nil
}
