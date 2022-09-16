package driver

import (
	controller "go-gameroom/adapter/controller/web"
	"go-gameroom/adapter/gateway"
	presenter "go-gameroom/adapter/presenter/web"
	"go-gameroom/usecase/interactor"
	"log"
	"net/http"
)

func Serve(addr string) {
	game := controller.GameController{
		InputFactory:      interactor.NewGameInputPort,
		OutputFactory:     presenter.NewGameOutputPort,
		RepositoryFactory: gateway.NewGameRedisRepository,
	}
	http.HandleFunc("/game/", game.GetGameHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
