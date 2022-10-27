package route

import (
	"final_project/config"
	"final_project/urls"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	r := gin.Default()

	r.POST(urls.POST_USERS_LOGIN)
	r.POST(urls.POST_USERS_REGISTER)
	r.PUT(urls.PUT_USER)
	r.DELETE(urls.DELETE_USER)

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
