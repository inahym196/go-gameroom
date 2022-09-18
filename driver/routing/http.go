package routing

import (
	controller "go-gameroom/adapter/controller/http"
	gateway "go-gameroom/adapter/gateway/inmemory"
	presenter "go-gameroom/adapter/presenter/http"
	"go-gameroom/usecase/interactor"
	"log"
	"net/http"
)

type Router interface {
	Run(addr string)
}

type HTTPRouter struct{}

func NewHTTPRouter() Router {
	return &HTTPRouter{}
}

func (r *HTTPRouter) Run(addr string) {
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

}
