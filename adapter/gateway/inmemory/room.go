package gateway

import (
	"fmt"
	"go-gameroom/domain/entity"
)

type RoomRepository struct{}

var InMemoryDataBase [4]*entity.Room = [4]*entity.Room{&entity.Room{Id: "0"}, &entity.Room{Id: "1"}, &entity.Room{Id: "2"}, &entity.Room{Id: "3"}}

func NewRoomRepository() entity.RoomRepository {
	return &RoomRepository{}
}

func (repo *RoomRepository) GetRoomById(roomId entity.RoomId) (*entity.Room, error) {
	fmt.Printf("%#v", roomId)
	i := roomId.MustInt()
	if !(0 <= i && i <= 4) {
		return nil, fmt.Errorf("Index Out of Range")
	}
	return InMemoryDataBase[i], nil
}

func (repo *RoomRepository) GetRooms() (map[int]*entity.Room, error) {
	repositoryFormatData := make(map[int]*entity.Room, 4)
	for i, room := range InMemoryDataBase {
		repositoryFormatData[i] = room
	}
	return repositoryFormatData, nil
}

func (repo *RoomRepository) Init(roomId entity.RoomId) (*entity.Room, error) {
	room := &entity.Room{
		Id: roomId,
	}
	i := roomId.MustInt()
	if !(0 <= i && i <= 4) {
		return nil, fmt.Errorf("Index Out of Range")
	}
	InMemoryDataBase[i] = room
	return room, nil
}
