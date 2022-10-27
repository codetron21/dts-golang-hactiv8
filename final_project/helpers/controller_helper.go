package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	contentType         = "Content-Type"
	ApplicationJsonType = "application/json"
)

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get(contentType)
}

func InCorrectContentTypeResponse() gin.H {
	return gin.H{
		"code":    http.StatusBadRequest,
		"message": "content type must application/json",
	}
}

func BindJsonFailResponse() gin.H {
	return gin.H{
		"code":    http.StatusInternalServerError,
		"message": "error bind json",
	}
}

func CreateErrorResponse(status int, message string) gin.H {
	return gin.H{
		"code":    status,
		"message": message,
	}
}
