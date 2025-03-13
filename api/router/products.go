package router

import (
	"indico-technical-test/app/controllers"
	"indico-technical-test/app/repository"
	"indico-technical-test/app/usecase"
	"indico-technical-test/database"

	"github.com/go-playground/validator/v10"
)

func Products() controllers.IProductController {
	db := database.Postgres{}
	validator := validator.New()
	repository := repository.NewProductRepository(db.Get())
	usecase := usecase.NewProductUsecase(repository)
	controllers := controllers.NewProductController(validator, usecase)
	return controllers
}
