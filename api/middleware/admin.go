package middleware

import (
	"indico-technical-test/app/entity"
	"indico-technical-test/database"
	"strings"

	"github.com/gin-gonic/gin"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		db := database.Postgres{}

		var user entity.User
		err := db.Get().GDB.Table("users").Where("token = ?", strings.ReplaceAll(token, "Bearer ", "")).First(&user).Error
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error"})
			c.Abort()
			return
		}

		if !strings.Contains(user.Roles, "admin") {
			c.JSON(403, gin.H{"message": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
