package driver

import (
	controller "go-gameroom/adapter/controller/http"
	gateway "go-gameroom/adapter/gateway/inmemory"
	presenter "go-gameroom/adapter/presenter/http"
	"go-gameroom/usecase/interactor"
	"log"
	"net/http"
)

func Serve2(addr string) {
	game := controller.GameController{
		InputFactory:      interactor.NewXOGameInputPort,
		OutputFactory:     presenter.NewXOGameOutputPort,
		RepositoryFactory: gateway.NewXOGameRepository,
	}
	http.HandleFunc("/game/", game.GetGameHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	// mrouter := routing.GetMRouter()
	// upgradeHandleFunc := routing.UpgradeHandleFunc(mrouter)
	// router := routing.GetRouter(upgradeHandleFunc)
	// router.Run(addr)
}
