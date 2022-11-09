package controller

import (
	"final_project/helpers"
	"final_project/middleware"
	"final_project/model"
	"final_project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CommentController struct {
	service service.CommentService
}

func (ctl CommentController) GetComments(ctx *gin.Context) {
	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))

	results, err := ctl.service.GetComments(userIdToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse(
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, results)
}

func (ctl CommentController) CreateComment(ctx *gin.Context) {
	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))
	comment := getCommentRequest(ctx)

	err := ctl.service.CreateComment(userIdToken, comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse(
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func (ctl CommentController) UpdateCommentById(ctx *gin.Context) {
	commentIdStr := ctx.Param("commentId")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			"error id not recognized",
		))
		return
	}

	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))
	comment := getCommentRequest(ctx)

	err = ctl.service.UpdateCommentById(userIdToken, commentId, comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	})
}

func (ctl CommentController) DeleteCommentById(ctx *gin.Context) {
	commentIdStr := ctx.Param("commentId")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			"error id not recognized",
		))
		return
	}

	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))

	err = ctl.service.DeleteCommentById(userIdToken, commentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}

func getCommentRequest(ctx *gin.Context) *model.Comment {
	contentType := helpers.GetContentType(ctx)
	if contentType != helpers.ApplicationJsonType {
		ctx.JSON(http.StatusBadRequest, helpers.InCorrectContentTypeResponse())
		return nil
	}

	comment := model.Comment{}

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.BindJsonFailResponse())
		return nil
	}

	return &comment
}
