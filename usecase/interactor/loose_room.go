package interactor

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"strconv"
)

type RoomLooseInteractor struct {
	Repository entity.RoomRepository
}

func NewRoomLoosePort(repository entity.RoomRepository) port.RoomLoosePort {
	return &RoomLooseInteractor{
		Repository: repository,
	}
}

func (i *RoomLooseInteractor) GetRoomById(roomId port.RoomId) (*port.Room, error) {
	s := strconv.Itoa(roomId)
	res, err := i.Repository.GetRoomById(entity.RoomId(s))
	if err != nil {
		return nil, err
	}
	return roomTransfer(res), nil
}

func (i *RoomLooseInteractor) GetRooms() (map[int]*port.Room, error) {
	res, err := i.Repository.GetRooms()
	if err != nil {
		return nil, err
	}
	rooms := make(map[int]*port.Room)
	for i, room := range res {
		if room.Id != "" {
			rooms[i] = roomTransfer(room)
		}
	}
	return rooms, nil
}

func (i *RoomLooseInteractor) Init(roomId port.RoomId) (*port.Room, error) {
	s := strconv.Itoa(roomId)
	res, err := i.Repository.Init(entity.RoomId(s))
	if err != nil {
		return nil, err
	}
	return roomTransfer(res), nil
}
