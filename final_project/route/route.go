package route

import (
	"final_project/config"
	"final_project/controller"
	"final_project/middleware"
	"final_project/urls"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(ctl controller.Controller) error {
	r := gin.Default()

	r.POST(urls.POST_USERS_LOGIN, ctl.UserController.Login)
	r.POST(urls.POST_USERS_REGISTER, ctl.UserController.Register)

	authorized := r.Group("/")
	{
		authorized.Use(middleware.Authentication())
		// user
		authorized.PUT(urls.PUT_USER, ctl.UserController.UpdateUserById)
		authorized.DELETE(urls.DELETE_USER, ctl.UserController.DeleteUserById)

		// photo
		authorized.POST(urls.POST_PHOTOS, ctl.PhotoController.CreatePhoto)
		authorized.GET(urls.GET_PHOTOS, ctl.PhotoController.GetPhotos)
		authorized.PUT(urls.PUT_PHOTO, ctl.PhotoController.UpdatePhotoById)
		authorized.DELETE(urls.DELETE_PHOTO, ctl.PhotoController.DeletePhotoById)

		// comment
		authorized.POST(urls.POST_COMMENTS)
		authorized.GET(urls.GET_COMMENTS)
		authorized.PUT(urls.PUT_COMMENT)
		authorized.DELETE(urls.DELETE_COMMENT)

		// social media
		authorized.POST(urls.POST_SOCIAL)
		authorized.GET(urls.GET_SOCIAL)
		authorized.PUT(urls.PUT_SOCIAL)
		authorized.DELETE(urls.DELETE_SOCIAL)
	}

	return r.Run(fmt.Sprintf("%s:%s", config.SERVER_HOST, config.SERVER_PORT))
}
