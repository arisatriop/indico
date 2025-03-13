package controllers

import (
	"fmt"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IOrderController interface {
	FindAll() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	Receive() gin.HandlerFunc
	Ship() gin.HandlerFunc
}

type OrderController struct {
	Validator      *validator.Validate
	OrderUsecase   usecase.IOrderUsecase
	ProductUsecase usecase.IProductUsecase
}

func NewOrderController(
	validator *validator.Validate,
	orderUsecase usecase.IOrderUsecase,
	productUsecase usecase.IProductUsecase,
) IOrderController {
	return &OrderController{
		Validator:      validator,
		OrderUsecase:   orderUsecase,
		ProductUsecase: productUsecase,
	}
}

func (c *OrderController) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		orders, err := c.OrderUsecase.FindAll()
		if err != nil {
			fmt.Printf("error find all orders: %v\n", err)
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    orders,
		})
	}
}

func (c *OrderController) FindByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ids := ctx.Param("id")
		id, _ := strconv.Atoi(ids)

		order, err := c.OrderUsecase.FindByID(id)
		if err != nil {
			fmt.Printf("error find order by id: %v\n", err)
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		if order == nil {
			ctx.JSON(404, gin.H{"message": "data not found"})
			return

		}

		ctx.JSON(200, gin.H{
			"message": "success",
			"data":    order,
		})
	}
}

func (c *OrderController) Receive() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var orders []request.OrderRequest
		if err := ctx.ShouldBindJSON(&orders); err != nil {
			ctx.JSON(400, gin.H{"message": "invalid request payload"})
			return
		}

		for _, order := range orders {
			if err := c.Validator.Struct(order); err != nil {
				ctx.JSON(400, gin.H{"message": err.Error()})
				return
			}
		}

		for _, order := range orders {
			product, err := c.ProductUsecase.FindByID(order.ProductID)
			if err != nil {
				fmt.Printf("error find product by id: %v\n", err)
				ctx.JSON(500, gin.H{"message": "internal server error"})
				return
			}
			if product == nil {
				ctx.JSON(404, gin.H{"message": "product not found"})
				return
			}
		}

		if err := c.OrderUsecase.Receive(orders); err != nil {
			fmt.Printf("error receive order: %v\n", err)
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(201, gin.H{"message": "success"})
	}
}

func (c *OrderController) Ship() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var orders []request.OrderRequest
		if err := ctx.ShouldBindJSON(&orders); err != nil {
			ctx.JSON(400, gin.H{"message": "invalid request payload"})
			return
		}

		for _, order := range orders {
			if err := c.Validator.Struct(order); err != nil {
				ctx.JSON(400, gin.H{"message": err.Error()})
				return
			}
		}

		for _, order := range orders {
			product, err := c.ProductUsecase.FindByID(order.ProductID)
			if err != nil {
				fmt.Printf("error find product by id: %v\n", err)
				ctx.JSON(500, gin.H{"message": "internal server error"})
				return
			}
			if product == nil {
				ctx.JSON(404, gin.H{"message": "product not found"})
				return
			}
			if product.Quantity < order.Quantity {
				ctx.JSON(400, gin.H{"message": "insufficient stock"})
				return
			}
		}

		if err := c.OrderUsecase.Ship(orders); err != nil {
			fmt.Printf("error receive order: %v\n", err)
			ctx.JSON(500, gin.H{"message": "internal server error"})
			return
		}

		ctx.JSON(201, gin.H{"message": "success"})
	}
}
