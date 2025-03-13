package main

import (
	"indico-technical-test/api/router"
	"indico-technical-test/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load(".env")
}

func main() {

	err := database.NewPostgres().Set()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := gin.Default()
	r.Use(CustomRecoveryMiddleware())

	router.Init(r)

	r.Run("localhost:8080")

}

func CustomRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				log.Printf("Recovered from panic: %v", err)

				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server error",
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
