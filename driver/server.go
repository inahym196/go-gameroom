package driver

import (
	controller "go-gameroom/adapter/controller/gin"
	gateway "go-gameroom/adapter/gateway/inmemory"
	"go-gameroom/driver/routing"
	"go-gameroom/usecase/interactor"
)

func Serve(addr string) {
	roomController := controller.RoomController{
		RepositoryFactory: gateway.NewRoomRepository,
		InputPortFactory:  interactor.NewRoomInputPort,
	}
	userController := controller.UserController{
		UserRepositoryFactory:    gateway.NewUserRepository,
		SessionRepositoryFactory: gateway.NewSessionRepository,
		InputPortFactory:         interactor.NewUserInputPort,
	}
	// router := routing.NewHTTPRouter()
	router := routing.NewGinRouter(roomController, userController)
	router.Run(addr)
}
