package router

import (
	"indico-technical-test/app/controllers"
	"indico-technical-test/app/repository"
	"indico-technical-test/app/usecase"
	"indico-technical-test/database"
)

func Users() controllers.IUserController {
	db := &database.Postgres{}
	repository := repository.NewUserRepository(db.Get())
	usecase := usecase.NewUserUsecase(repository)
	controllers := controllers.NewUserController(usecase)
	return controllers
}
