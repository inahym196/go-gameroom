package driver

import (
	controller "go-gameroom/adapter/controller/http"
	gateway "go-gameroom/adapter/gateway/inmemory"
	presenter "go-gameroom/adapter/presenter/http"
	"go-gameroom/usecase/interactor"
	"log"
	"net/http"
)

func Serve(addr string) {
	room := controller.RoomController{
		InputFactory:      interactor.NewRoomInputPort,
		OutputFactory:     presenter.NewRoomOutputPort,
		RepositoryFactory: gateway.NewRoomRepository,
	}
	http.HandleFunc("/rooms/", room.EndpointHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	//mrouter := routing.GetMRouter()
	//upgradeHandleFunc := routing.UpgradeHandleFunc(mrouter)
	//router := routing.GetRouter(upgradeHandleFunc)
	//router.Run(addr)
}
