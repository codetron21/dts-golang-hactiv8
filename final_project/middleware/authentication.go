package middleware

import (
	"final_project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

const USER_DATA = "user-data"

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized,
				helpers.CreateErrorResponse(
					http.StatusUnauthorized,
					err.Error(),
				),
			)
			return
		}

		ctx.Set(USER_DATA, verifyToken)
		ctx.Next()
	}
}
