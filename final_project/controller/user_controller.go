package controller

import (
	"final_project/helpers"
	"final_project/middleware"
	"final_project/model"
	"final_project/service"
	"final_project/urls"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserController struct {
	service service.UserService
}

func (ctl UserController) Register(ctx *gin.Context) {
	user := getUserRequest(ctx)

	err := ctl.service.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			helpers.CreateErrorResponse(
				http.StatusInternalServerError,
				"error create user data",
			))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})
}

func (ctl UserController) Login(ctx *gin.Context) {
	user := getUserRequest(ctx)

	token, err := ctl.service.LoginUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"token": token,
		},
	)
}

func (ctl UserController) DeleteUserById(ctx *gin.Context) {
	userIdStr := ctx.Param(urls.PATH_USER_ID)

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			"Invalid path user id",
		))
		return
	}

	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))

	err = ctl.service.DeleteUserById(userIdToken, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}

func (ctl UserController) UpdateUserById(ctx *gin.Context) {
	newUser := getUserRequest(ctx)

	userIdStr := ctx.Param(urls.PATH_USER_ID)

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			"Invalid path user id",
		))
		return
	}

	userData := ctx.MustGet(middleware.USER_DATA).(jwt.MapClaims)
	userIdToken := int(userData[helpers.CLAIM_ID].(float64))

	updatedUser, err := ctl.service.UpdateUserById(newUser, userIdToken, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse(
			http.StatusBadRequest,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         updatedUser.ID,
		"email":      updatedUser.Email,
		"username":   updatedUser.Username,
		"age":        updatedUser.Age,
		"updated_at": updatedUser.UpdatedAt,
	})
}

func getUserRequest(ctx *gin.Context) *model.User {
	contentType := helpers.GetContentType(ctx)
	if contentType != helpers.ApplicationJsonType {
		ctx.JSON(http.StatusBadRequest, helpers.InCorrectContentTypeResponse())
		return nil
	}

	user := model.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.BindJsonFailResponse())
		return nil
	}

	return &user
}
