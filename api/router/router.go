package router

import (
	m "indico-technical-test/api/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	api := r.Group("/api")

	api.POST("/register", Auth().Register())
	api.POST("/login", Auth().Login())

	users := api.Group("/users")
	users.GET("/me", m.Auth(), Users().FindMe())
	users.GET("/", m.Auth(), m.Admin(), Users().FindAll())

	locations := api.Group("/locations")
	locations.POST("/", m.Auth(), m.Admin(), Locations().Create())
	locations.GET("/", m.Auth(), Locations().FindAll())

	products := api.Group("/products")
	products.POST("/", m.Auth(), m.Admin(), Products().Create())
	products.GET("/", m.Auth(), Products().FindAll())
	products.GET("/:id", m.Auth(), Products().FindByID())
	products.PUT("/:id", m.Auth(), m.Admin(), Products().Update())
	products.DELETE("/:id", m.Auth(), m.Admin(), Products().Delete())

	orders := api.Group("/orders")
	orders.POST("/receive", m.Auth(), m.Staff(), Orders().Receive())
	orders.POST("/ship", m.Auth(), m.Staff(), Orders().Ship())
	orders.GET("/", m.Auth(), Orders().FindAll())
	orders.GET("/:id", m.Auth(), Orders().FindByID())

}
