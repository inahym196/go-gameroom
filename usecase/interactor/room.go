package interactor

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
)

type RoomInteractor struct {
	OutputPort port.RoomOutputPort
	Repository entity.RoomRepository
}

func NewRoomInputPort(outputPort port.RoomOutputPort, repository entity.RoomRepository) port.RoomInputPort {
	return &RoomInteractor{
		OutputPort: outputPort,
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

func (i *RoomInteractor) GetRoomById(roomId port.RoomId) {
	res, err := i.Repository.GetRoomById(entity.RoomId(roomId))
	if err != nil {
		i.OutputPort.GetRoomById(nil, err)
		return
	}
	outputData := roomTransfer(res)
	i.OutputPort.GetRoomById(outputData, nil)
}

func (i *RoomInteractor) GetRooms() {
	res, err := i.Repository.GetRooms()
	if err != nil {
		i.OutputPort.GetRooms(nil, err)
	}
	rooms := make(map[int]*port.Room)
	for i, room := range res {
		if room.Id != "" {
			rooms[i] = roomTransfer(room)
		}
	}
	i.OutputPort.GetRooms(rooms, nil)
}

func (i *RoomInteractor) Init(roomId port.RoomId) {
	res, err := i.Repository.Init(entity.RoomId(roomId))
	if err != nil {
		i.OutputPort.Init(nil, err)
		return
	}
	outputData := roomTransfer(res)
	i.OutputPort.Init(outputData, nil)
}
