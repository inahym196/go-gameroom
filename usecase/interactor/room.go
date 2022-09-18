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

func roomTransfer(entityRoom *entity.Room) *port.Room {
	var atendee []string
	for _, user := range entityRoom.Atendee {
		atendee = append(atendee, user.Name)
	}
	outputRoom := &port.Room{
		RoomId:     entityRoom.Id.MustInt(),
		Url:        string(entityRoom.Url),
		Atendee:    atendee,
		GameStatus: string(entityRoom.Game.Status),
	}
	return outputRoom
}

func (i *RoomInteractor) GetRoomById(roomId port.RoomId) (*port.Room, error) {
	s := strconv.Itoa(roomId)
	res, err := i.Repository.GetRoomById(entity.RoomId(s))
	if err != nil {
		return nil, err
	}
	return roomTransfer(res), nil
}

func (i *RoomInteractor) GetRooms() (map[int]*port.Room, error) {
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

func (i *RoomInteractor) Init(roomId port.RoomId) (*port.Room, error) {
	s := strconv.Itoa(roomId)
	res, err := i.Repository.Init(entity.RoomId(s))
	if err != nil {
		return nil, err
	}
	return roomTransfer(res), nil
}
