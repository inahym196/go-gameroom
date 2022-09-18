package controller

import (
	"fmt"
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"net/http"
	"strings"
)

type RoomController struct {
	OutputFactory     func(w http.ResponseWriter) port.RoomOutputPort
	RepositoryFactory func() entity.RoomRepository
	InputFactory      func(outputPort port.RoomOutputPort, repository entity.RoomRepository) port.RoomInputPort
}

func (c *RoomController) EndpointHandler(w http.ResponseWriter, r *http.Request) {
	outputPort := c.OutputFactory(w)
	repository := c.RepositoryFactory()
	inputPort := c.InputFactory(outputPort, repository)
	roomId := strings.TrimPrefix(r.URL.Path, "/rooms/")

	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		if len(roomId) > 0 {
			inputPort.GetRoomById(roomId)
		} else {
			inputPort.GetRooms()
		}
	case http.MethodPost:
		fmt.Println("POST")
		inputPort.Create(roomId)
	case http.MethodDelete:
		fmt.Println("DELETE")
		inputPort.Delete(roomId)
	}
}
