package port

import (
	"go-gameroom/domain/entity"
	"strconv"
)

type RoomDto struct {
	Id         int
	Url        string
	Atendee    []string
	GameStatus string
}

type RoomInputPort interface {
	GetRoomById(Id int) (*RoomDto, error)
	GetRooms() (map[string]*RoomDto, error)
	Init(Id int) (*RoomDto, error)
}

func NewRoomDtoFromEntity(room *entity.Room) *RoomDto {
	roomId, _ := strconv.Atoi(room.Id)
	var roomAtendee []string
	for _, user := range room.Atendee {
		roomAtendee = append(roomAtendee, user.Name)
	}
	return &RoomDto{
		Id:         roomId,
		Url:        room.Url,
		Atendee:    roomAtendee,
		GameStatus: room.Game.Status(),
	}
}
