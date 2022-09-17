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

func roomTransfer(entityRoom *entity.Room) (outputRoom *port.Room) {
	var atendee []string
	for _, user := range entityRoom.Atendee {
		atendee = append(atendee, user.Name)
	}
	outputRoom = &port.Room{
		RoomId:     string(entityRoom.Id),
		Url:        string(entityRoom.Url),
		Atendee:    atendee,
		GameStatus: string(entityRoom.Game.Status),
	}
	return outputRoom

}

func (i *RoomInteractor) GetRoomById(roomId port.RoomId) {
	res, err := i.Repository.GetRoomById(entity.RoomId(roomId))
	if err != nil {
		panic(0)
	}
	outputData := roomTransfer(res)
	i.OutputPort.GetRoomById(outputData)
}

func (i *RoomInteractor) GetRooms() {
	res, err := i.Repository.GetRooms()
	if err != nil {
		panic(0)
	}
	var rooms []*port.Room
	for _, room := range res {
		rooms = append(rooms, roomTransfer(room))
	}
	i.OutputPort.GetRooms(rooms)
}
