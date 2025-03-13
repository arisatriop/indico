package controllers

import (
	"indico-technical-test/app/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	FindAll() gin.HandlerFunc
	FindMe() gin.HandlerFunc
}

type UserController struct {
	UserUseCase usecase.IUserUsecase
}

func NewUserController(userUsecase usecase.IUserUsecase) IUserController {
	return &UserController{
		UserUseCase: userUsecase,
	}
}

func (c *UserController) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := c.UserUseCase.FindAll()
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Internal server error"})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    users,
		})
	}
}

func (c *UserController) FindMe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := c.UserUseCase.FindMe(strings.ReplaceAll(ctx.GetHeader("Authorization"), "Bearer ", ""))
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Internal server error"})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    user,
		})
	}
}
