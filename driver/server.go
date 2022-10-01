package driver

import (
	controller "go-gameroom/adapter/controller/gin"
	gateway "go-gameroom/adapter/gateway/inmemory"
	"go-gameroom/driver/routing"
	"go-gameroom/usecase/interactor"
)

func Serve(addr string) {
	userController := controller.UserController{
		UserRepositoryFactory:    gateway.NewUserRepository,
		SessionRepositoryFactory: gateway.NewSessionRepository,
		InputPortFactory:         interactor.NewUserInputPort,
	}
	gameController := controller.GameController{
		RepositoryFactory: gateway.NewXOGameRepository,
		InputPortFactory:  interactor.NewXOGameInputPort,
	}
	// router := routing.NewHTTPRouter()
	router := routing.NewGinRouter(userController, gameController)
	router.Run(addr)
}
