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

type PhotoController struct {
	service service.PhotoService
}

func (ctl PhotoController) CreatePhoto(ctx *gin.Context) {
	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))

	photoData := getPhotoRequest(ctx)

	err := ctl.service.CreatePhoto(userIdToken, photoData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse(
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         photoData.ID,
		"title":      photoData.Title,
		"caption":    photoData.Caption,
		"photo_url":  photoData.PhotoUrl,
		"user_id":    photoData.UserID,
		"created_at": photoData.CreatedAt,
	})
}

func (ctl PhotoController) GetPhotos(ctx *gin.Context) {
	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))

	results, err := ctl.service.GetPhotos(userIdToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse(
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, results)
}

func (ctl PhotoController) UpdatePhotoById(ctx *gin.Context) {
	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))
	photoData := getPhotoRequest(ctx)
	photoId := ctx.Param("photoId")

	value, err := strconv.Atoi(photoId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			"error id not recognized",
		))
		return
	}

	err = ctl.service.UpdatePhotoById(photoData, value, userIdToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse(
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         photoData.ID,
		"title":      photoData.Title,
		"caption":    photoData.Caption,
		"photo_url":  photoData.PhotoUrl,
		"user_id":    photoData.UserID,
		"updated_at": photoData.UpdatedAt,
	})
}

func (ctl PhotoController) DeletePhotoById(ctx *gin.Context) {
	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))
	photoId := ctx.Param("photoId")

	value, err := strconv.Atoi(photoId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			"error id not recognized",
		))
		return
	}

	err = ctl.service.DeletePhotoById(value, userIdToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse(
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}

func getPhotoRequest(ctx *gin.Context) *model.Photo {
	contentType := helpers.GetContentType(ctx)
	if contentType != helpers.ApplicationJsonType {
		ctx.JSON(http.StatusBadRequest, helpers.InCorrectContentTypeResponse())
		return nil
	}

	photo := model.Photo{}

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.BindJsonFailResponse())
		return nil
	}

	return &photo
}
