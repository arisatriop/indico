package controllers

import (
	"indico-technical-test/app/entity/request"
	"indico-technical-test/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ILocationController interface {
	Create() gin.HandlerFunc
	FindAll() gin.HandlerFunc
}

type LocationController struct {
	Validator       *validator.Validate
	LocationUsecase usecase.ILocationUsecase
}

func NewLocationController(validator *validator.Validate, locationUsecase usecase.ILocationUsecase) ILocationController {
	return &LocationController{
		Validator:       validator,
		LocationUsecase: locationUsecase,
	}
}

func (c *LocationController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request.LocationRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(400, gin.H{"message": "invalid request payload"})
			return
		}

		if err := c.Validator.Struct(req); err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			return
		}

		if err := c.LocationUsecase.Create(&req); err != nil {
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(201, gin.H{"message": "success"})
	}
}

func (c *LocationController) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		locations, err := c.LocationUsecase.FindAll()
		if err != nil {
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    locations,
		})
	}
}
