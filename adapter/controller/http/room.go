package controller

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"net/http"
)

type RoomController struct {
	OutputFactory     func(w http.ResponseWriter) port.RoomOutputPort
	RepositoryFactory func() entity.RoomRepository
	InputFactory      func(outputPort port.RoomOutputPort, repository entity.RoomRepository) port.RoomInputPort
}

func (c *RoomController) GetRooms(w http.ResponseWriter, r *http.Request) {
	outputPort := c.OutputFactory(w)
	repository := c.RepositoryFactory()
	inputPort := c.InputFactory(outputPort, repository)
	inputPort.GetRooms()
}
