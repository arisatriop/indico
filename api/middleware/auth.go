package middleware

import (
	"indico-technical-test/app/entity"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		claims := &entity.Claims{}
		token, err := jwt.ParseWithClaims(strings.ReplaceAll(tokenStr, "Bearer ", ""), claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Token expired"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			}
			c.Abort()
			return
		}

		c.Next()
	}
}
