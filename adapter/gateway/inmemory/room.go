package gateway

import (
	"fmt"
	"go-gameroom/domain/entity"
)

type RoomRepository struct{}

var InMemoryDataBase []*entity.Room

func NewRoomRepository() entity.RoomRepository {
	return &RoomRepository{}
}

func (repo *RoomRepository) GetRoomById(roomId entity.RoomId) (*entity.Room, bool, error) {
	for _, room := range InMemoryDataBase {
		if string(room.Id) == string(roomId) {
			return room, true, nil
		}
	}
	return nil, false, fmt.Errorf("not found")
}

func (repo *RoomRepository) GetRooms() (map[int]*entity.Room, error) {
	repositoryFormatData := make(map[int]*entity.Room)
	for i, room := range InMemoryDataBase {
		repositoryFormatData[i] = room
	}
	return repositoryFormatData, nil
}

func (repo *RoomRepository) Create(roomId entity.RoomId) (*entity.Room, bool) {
	for _, room := range InMemoryDataBase {
		if string(room.Id) == string(roomId) {
			return nil, false
		}
	}
	room := &entity.Room{
		Id: roomId,
	}
	InMemoryDataBase = append(InMemoryDataBase, room)
	return room, true
}

func (repo *RoomRepository) Delete(roomId entity.RoomId) (bool, error) {
	fmt.Println(len(InMemoryDataBase))
	for i, room := range InMemoryDataBase {
		if string(room.Id) == string(roomId) {
			InMemoryDataBase[i] = InMemoryDataBase[len(InMemoryDataBase)-1]
			InMemoryDataBase = InMemoryDataBase[:len(InMemoryDataBase)-1]
			return true, nil
		}
	}
	return false, nil
}
