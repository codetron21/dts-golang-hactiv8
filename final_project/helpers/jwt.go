package helpers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const secretKey = "jfdalfjdlajfdjafdjalfjdalkfnv"
const CLAIM_ID = "id"

func GenerateToken(id int, email string) (string, error) {
	claims := jwt.MapClaims{
		CLAIM_ID: id,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if headerToken == "" {
		return nil, errors.New("authorization header must be included")
	}

	if !bearer {
		return nil, errors.New("auth token must use bearer keyword")
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error parsing token")
		}

		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New("error claims token")
	}

	return token.Claims.(jwt.MapClaims), nil
}
