package gateway

import (
	"fmt"
	"go-gameroom/domain/entity"
	"strconv"
)

type RoomRepository struct{}

var RoomDataBase [4]*entity.Room = [4]*entity.Room{&entity.Room{Id: "0"}, &entity.Room{Id: "1"}, &entity.Room{Id: "2"}, &entity.Room{Id: "3"}}

func NewRoomRepository() entity.RoomRepository {
	return &RoomRepository{}
}

func (repo *RoomRepository) GetRoomById(roomId string) (*entity.Room, error) {
	fmt.Printf("%#v", roomId)
	i, _ := strconv.Atoi(roomId)
	if !(0 <= i && i <= 4) {
		return nil, fmt.Errorf("Index Out of Range")
	}
	return RoomDataBase[i], nil
}

func (repo *RoomRepository) GetRooms() (map[string]*entity.Room, error) {
	repositoryFormatData := make(map[string]*entity.Room, 4)
	for i, room := range RoomDataBase {
		s := strconv.Itoa(i)
		repositoryFormatData[s] = room
	}
	return repositoryFormatData, nil
}

func (repo *RoomRepository) Init(roomId string) (*entity.Room, error) {
	room := &entity.Room{
		Id: roomId,
	}
	i, _ := strconv.Atoi(roomId)
	if !(0 <= i && i <= 4) {
		return nil, fmt.Errorf("Index Out of Range")
	}
	RoomDataBase[i] = room
	return room, nil
}
