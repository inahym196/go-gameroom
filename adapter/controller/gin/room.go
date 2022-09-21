package controller

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
)

type RoomController struct {
	RepositoryFactory func() entity.RoomRepository
	InputPortFactory  func(repository entity.RoomRepository) port.RoomInputPort
}

func (c *RoomController) GetRooms() (map[string]*port.RoomDto, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	rooms, err := inputport.GetRooms()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
