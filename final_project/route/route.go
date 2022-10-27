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
		authorized.PUT(urls.PUT_USER, ctl.UserController.UpdateUserById)
		authorized.DELETE(urls.DELETE_USER, ctl.UserController.DeleteUserById)
	}

	r.POST(urls.POST_PHOTOS)
	r.GET(urls.GET_PHOTOS)
	r.PUT(urls.PUT_PHOTO)
	r.DELETE(urls.DELETE_PHOTO)

	r.POST(urls.POST_COMMENTS)
	r.GET(urls.GET_COMMENTS)
	r.PUT(urls.PUT_COMMENT)
	r.DELETE(urls.DELETE_COMMENT)

	r.POST(urls.POST_SOCIAL)
	r.GET(urls.GET_SOCIAL)
	r.PUT(urls.PUT_SOCIAL)
	r.DELETE(urls.DELETE_SOCIAL)

	return r.Run(fmt.Sprintf("%s:%s", config.SERVER_HOST, config.SERVER_PORT))
}
