package interactor

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"strconv"
)

type RoomInteractor struct {
	Repository entity.RoomRepository
}

func NewRoomInputPort(repository entity.RoomRepository) port.RoomInputPort {
	return &RoomInteractor{
		Repository: repository,
	}
}

func (i *RoomInteractor) GetRoomById(roomId int) (*port.RoomDto, error) {
	s := strconv.Itoa(roomId)
	room, err := i.Repository.GetRoomById(s)
	if err != nil {
		return nil, err
	}
	return port.NewRoomDtoFromEntity(room), nil
}

func (i *RoomInteractor) GetRooms() (map[string]*port.RoomDto, error) {
	res, err := i.Repository.GetRooms()
	if err != nil {
		return nil, err
	}
	rooms := make(map[string]*port.RoomDto)
	for i, room := range res {
		if room.Id != "" {
			rooms[i] = port.NewRoomDtoFromEntity(room)
		}
	}
	return rooms, nil
}

func (i *RoomInteractor) Init(roomId int) (*port.RoomDto, error) {
	s := strconv.Itoa(roomId)
	room, err := i.Repository.Init(s)
	if err != nil {
		return nil, err
	}
	return port.NewRoomDtoFromEntity(room), nil
}
