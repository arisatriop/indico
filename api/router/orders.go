package router

import (
	"indico-technical-test/app/controllers"
	"indico-technical-test/app/repository"
	"indico-technical-test/app/usecase"
	"indico-technical-test/database"

	"github.com/go-playground/validator/v10"
)

func Orders() controllers.IOrderController {
	db := database.Postgres{}
	pg := db.Get()

	validator := validator.New()
	productRepository := repository.NewProductRepository(pg)
	productUsecase := usecase.NewProductUsecase(productRepository)
	repository := repository.NewOrderRepository(pg)
	usecase := usecase.NewOrderUsecase(pg, repository, productRepository)
	controllers := controllers.NewOrderController(validator, usecase, productUsecase)
	return controllers
}
