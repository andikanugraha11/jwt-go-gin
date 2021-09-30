package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
)

var secret = "shhhh!"

func GenerateToken(id uint, email string) string{
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secret))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error){
	errResponse := errors.New("please sigint to process")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	// Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
	// Authorization: Basic username:password
	if !bearer {
		return nil, errResponse
	}

	// ["Bearer","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"]
	stringToken := strings.Split(headerToken," ")[1]

	// verify process
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (i interface{}, err error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secret), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, errResponse
	}

	if !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil

}