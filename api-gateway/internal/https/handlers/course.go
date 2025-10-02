package handlers

import (
	"api-gateway/genproto/coursepb"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HandlerSt) CreateCourse(ctx *gin.Context) {

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	defer file.Close()

	// Unique fayl nomi yaratish
	filename := header.Filename
	uniqueKey := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filename)

	result, err := h.CloudFlare.UploadFile(ctx, uniqueKey, file)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	req := coursepb.CreateCourseReq{
		UserId:       ctx.PostForm("user_id"),
		Title:        ctx.PostForm("title"),
		Description:  ctx.PostForm("description"),
		Category:     ctx.PostForm("category"),
		Level:        ctx.PostForm("level"),
		ThumbnailUrl: result.URL, // faylni yuklaganimizdan keyin kelgan URL
	}

	resp, err := h.Service.CreateCourse(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *HandlerSt) UpdateCourse(ctx *gin.Context) {
	req := coursepb.UpdateCourseReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = ctx.Param("id")
	resp, err := h.Service.UpdateCourse(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *HandlerSt) DeleteCourse(ctx *gin.Context) {
	req := coursepb.DeleteCourseReq{Id: ctx.Param("id")}
	resp, err := h.Service.DeleteCourse(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *HandlerSt) GetCourseById(ctx *gin.Context) {
	req := coursepb.GetCourseByIdReq{Id: ctx.Param("id")}
	resp, err := h.Service.GetCourseById(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *HandlerSt) GetCourses(ctx *gin.Context) {
	req := coursepb.GetCoursesReq{}
	resp, err := h.Service.GetCourses(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *HandlerSt) GetCourseByUserId(ctx *gin.Context) {
	req := coursepb.GetCourseByUserIdReq{UserId: ctx.Param("user_id")}
	resp, err := h.Service.GetCourseByUserId(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
