package controllers

import (
	"fmt"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IProductController interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	FindAll() gin.HandlerFunc
	FindByID() gin.HandlerFunc
}

type ProductController struct {
	Validator      *validator.Validate
	ProductUsecase usecase.IProductUsecase
}

func NewProductController(validator *validator.Validate, productUsecase usecase.IProductUsecase) IProductController {
	return &ProductController{
		Validator:      validator,
		ProductUsecase: productUsecase,
	}
}

func (c *ProductController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request.ProductCreateRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(400, gin.H{"message": "invalid request payload"})
			return
		}

		if err := c.Validator.Struct(req); err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			return
		}

		if err := c.ProductUsecase.Create(&req); err != nil {
			fmt.Printf("error create product: %v\n", err)
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(201, gin.H{"message": "success"})
	}
}

func (c *ProductController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request.ProductUpdateRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(400, gin.H{"message": "invalid request payload"})
			return
		}

		if err := c.Validator.Struct(req); err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			return
		}

		if err := c.ProductUsecase.Update(&req); err != nil {
			fmt.Printf("error update product: %v\n", err)
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "success"})
	}
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ids := ctx.Param("id")
		id, _ := strconv.Atoi(ids)

		product, err := c.ProductUsecase.FindByID(id)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		if product == nil {
			ctx.JSON(404, gin.H{"message": "product not found"})
			return
		}

		if err := c.ProductUsecase.Delete(id); err != nil {
			fmt.Printf("error delete product: %v\n", err)
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "success"})
	}
}

func (c *ProductController) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		products, err := c.ProductUsecase.FindAll()
		if err != nil {
			fmt.Printf("error find all product: %v\n", err)
			return
		}

		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    products,
		})
	}
}

func (c *ProductController) FindByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ids := ctx.Param("id")
		id, _ := strconv.Atoi(ids)

		product, err := c.ProductUsecase.FindByID(id)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		if product == nil {
			ctx.JSON(404, gin.H{"message": "product not found"})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    product,
		})
	}
}
