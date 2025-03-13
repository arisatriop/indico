package router

import (
	"indico-technical-test/app/controllers"
	"indico-technical-test/app/repository"
	"indico-technical-test/app/usecase"
	"indico-technical-test/database"

	"github.com/go-playground/validator/v10"
)

func Auth() controllers.IAuthController {
	db := &database.Postgres{}
	validator := validator.New()
	repository := repository.NewAuthRepository(db.Get())
	usecase := usecase.NewAuthUsecase(repository)
	controllers := controllers.NewAuthController(validator, usecase)
	return controllers
}
