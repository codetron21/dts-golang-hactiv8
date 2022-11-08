package controller

import (
	"final_project/service"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service service.CommentService
}

func (ctl CommentController) GetComments(ctx *gin.Context) {

}

func (ctl CommentController) CreateComment(ctx *gin.Context) {

}

func (ctl CommentController) UpdateCommentById(ctx *gin.Context) {

}

func (ctl CommentController) DeleteCommentById(ctx *gin.Context) {

}
