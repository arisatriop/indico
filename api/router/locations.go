package router

import (
	"indico-technical-test/app/controllers"
	"indico-technical-test/app/repository"
	"indico-technical-test/app/usecase"
	"indico-technical-test/database"

	"github.com/go-playground/validator/v10"
)

func Locations() controllers.ILocationController {
	db := &database.Postgres{}
	validator := validator.New()
	repository := repository.NewLocationRepository(db.Get())
	usecase := usecase.NewLocationUsecase(repository)
	controllers := controllers.NewLocationController(validator, usecase)
	return controllers
}
