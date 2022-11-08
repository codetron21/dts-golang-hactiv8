package route

import (
	"final_project/config"
	"final_project/controller"
	"final_project/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

const pathPhotoId = "photoId"
const pathCommentId = "commentId"
const pathSocialId = "socialMediaId"

func StartServer(ctl controller.Controller) error {
	r := gin.Default()

	userRoute := r.Group("/users")
	{
		userRoute.POST("/login", ctl.UserController.Login)
		userRoute.POST("/register", ctl.UserController.Register)
		userRoute.Use(middleware.Authentication())
		userRoute.PUT("", ctl.UserController.UpdateUserById)
		userRoute.DELETE("", ctl.UserController.DeleteUserById)
	}

	photoRoute := r.Group("/photos")
	{
		pathId := fmt.Sprintf("/:%s", pathPhotoId)

		photoRoute.Use(middleware.Authentication())

		photoRoute.POST("", ctl.PhotoController.CreatePhoto)
		photoRoute.GET("", ctl.PhotoController.GetPhotos)
		photoRoute.PUT(pathId, ctl.PhotoController.UpdatePhotoById)
		photoRoute.DELETE(pathId, ctl.PhotoController.DeletePhotoById)
	}

	commentRoute := r.Group("/comments")
	{
		pathId := fmt.Sprintf("/:%s", pathCommentId)

		commentRoute.Use(middleware.Authentication())

		commentRoute.POST("", ctl.CommentController.CreateComment)
		commentRoute.GET("", ctl.CommentController.GetComments)
		commentRoute.PUT(pathId, ctl.CommentController.UpdateCommentById)
		commentRoute.DELETE(pathId, ctl.CommentController.DeleteCommentById)
	}

	socialRoute := r.Group("/socialmedias")
	{
		pathId := fmt.Sprintf("/:%s", pathSocialId)

		socialRoute.Use(middleware.Authentication())

		socialRoute.POST("", ctl.SocialMediaController.CreateSocialMedia)
		socialRoute.GET("", ctl.SocialMediaController.GetSocialMedias)
		socialRoute.PUT(pathId, ctl.SocialMediaController.UpdateSocialMediaById)
		socialRoute.DELETE(pathId, ctl.SocialMediaController.DeleteSocialMediaById)
	}

	return r.Run(fmt.Sprintf("%s:%s", config.SERVER_HOST, config.SERVER_PORT))
}
