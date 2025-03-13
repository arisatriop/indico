package controllers

import (
	"fmt"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IAuthController interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type AuthController struct {
	Validator   *validator.Validate
	AuthUsecase usecase.IAuthUsecase
}

func NewAuthController(validator *validator.Validate, authUsecase usecase.IAuthUsecase) IAuthController {
	return &AuthController{
		Validator:   validator,
		AuthUsecase: authUsecase,
	}
}

func (c *AuthController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request.RegisterRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(400, gin.H{"message": "invalid request payload"})
			return
		}

		if err := c.Validator.Struct(req); err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			return
		}

		if err := c.AuthUsecase.Register(&req); err != nil {
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(201, gin.H{"message": "success"})
	}
}

func (c *AuthController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request.LoginRequest
		if err := ctx.ShouldBind(&req); err != nil {
			fmt.Printf("error binding json: %v", err)
			ctx.JSON(400, gin.H{"message": "invalid request payload"})
			return
		}

		if err := c.Validator.Struct(req); err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			return
		}

		token, err := c.AuthUsecase.Login(&req)
		if err != nil {
			fmt.Printf("error logging in: %v", err)
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(201, gin.H{
			"message": "success",
			"token":   token,
		})
	}
}
