package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nanwp/api-sederhana/config"
)

var Username string

func JWTMiddleware(c *gin.Context) {
	tokenstring, err := c.Cookie("token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unautorized 1",
		})
		c.Abort()
		return
	}

	claims := &config.JWTClaim{}

	token, err := jwt.ParseWithClaims(
		tokenstring,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		},
	)

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorSignatureInvalid:
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorize 2",
			})
			c.Abort()
			return
		case jwt.ValidationErrorExpired:
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorize! Token Expired",
			})
			c.Abort()
			return

		default:
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorize 3",
			})
			c.Abort()
			return
		}
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorize 4",
		})
		c.Abort()
		return
	}

	Username = claims.Username
	c.Next()
}
