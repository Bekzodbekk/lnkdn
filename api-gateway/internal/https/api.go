package https

import (
	"api-gateway/internal/https/handlers"
	"api-gateway/internal/https/middleware"
	cloudflareconnection "api-gateway/internal/pkg/CloudFlareConnection"
	"api-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

func Newgin(
	srv *service.ServiceRepositoryClient,
	cloudFlare *cloudflareconnection.CloudflareStorage,
) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.MiddleWare())

	hnd := handlers.NewHandlerSt(srv, cloudFlare)

	// Auth group
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/sign-in", hnd.SignIn)
		auth.POST("/forgot-password/send-code", hnd.ForgetPasswordSendCodeEmail)
		auth.POST("/forgot-password/check-code", hnd.ForgetPasswordCheckCode)
		auth.POST("/forgot-password/update", hnd.ForgetPasswordUpdatePassword)
		auth.POST("/update-password", hnd.UpdatePassword)
		auth.POST("/check-code-after-update", hnd.CheckCodeAfterUpdatePassword)
	}

	// Users group
	users := r.Group("/api/v1/users")
	{
		users.POST("/", hnd.CreateUser)
		users.PUT("/:id", hnd.UpdateUser)
		users.DELETE("/:id", hnd.DeleteUser)
		users.GET("/:id", hnd.GetUserById)
		users.GET("/", hnd.GetUsers)
	}

	// Posts group
	posts := r.Group("/api/v1/posts")
	{
		posts.POST("/", hnd.CreatePost)                      // post yaratish
		posts.PUT("/:id", hnd.UpdatePost)                    // id bo‘yicha yangilash
		posts.DELETE("/:id", hnd.DeletePost)                 // id bo‘yicha o‘chirish
		posts.GET("/:id", hnd.GetPostById)                   // id bo‘yicha olish
		posts.GET("/", hnd.GetAllPosts)                      // barcha postlarni olish
		posts.GET("/user/:user_id", hnd.GetAllPostsByUserId) // user_id bo‘yicha postlar
	}

	// Comments group
	comments := r.Group("/api/v1/comments")
	{
		comments.POST("/", hnd.CreateComment)                  // comment yaratish
		comments.PUT("/:id", hnd.UpdateComment)                // id bo‘yicha yangilash
		comments.DELETE("/:id", hnd.DeleteComment)             // id bo‘yicha o‘chirish
		comments.GET("/user/:user_id", hnd.GetCommentByUserId) // user_id bo‘yicha olish
		comments.GET("/post/:post_id", hnd.GetCommentByPostId) // post_id bo‘yicha olish
	}

	// Likes group
	likes := r.Group("/api/v1/likes")
	{
		likes.POST("/toggle", hnd.ToggleLike) // like/unlike (toggle)
	}

	// Courses group
	courses := r.Group("/api/v1/courses")
	{
		courses.POST("/", hnd.CreateCourse)                  // course yaratish
		courses.PUT("/:id", hnd.UpdateCourse)                // id bo‘yicha yangilash
		courses.DELETE("/:id", hnd.DeleteCourse)             // id bo‘yicha o‘chirish
		courses.GET("/:id", hnd.GetCourseById)               // id bo‘yicha olish
		courses.GET("/", hnd.GetCourses)                     // barcha course lar
		courses.GET("/user/:user_id", hnd.GetCourseByUserId) // user_id bo‘yicha course lar
	}

	// Lessons group
	// lessons := r.Group("/api/v1/lessons")
	// {
	// 	lessons.POST("/", hnd.CreateLesson)                         // lesson yaratish
	// 	lessons.PUT("/:id", hnd.UpdateLesson)                       // id bo‘yicha yangilash
	// 	lessons.DELETE("/:id", hnd.DeleteLesson)                    // id bo‘yicha o‘chirish
	// 	lessons.GET("/:id", hnd.GetLessonById)                      // id bo‘yicha olish
	// 	lessons.GET("/course/:course_id", hnd.GetLessonsByCourseId) // course_id bo‘yicha lessonlar
	// }

	return r
}
