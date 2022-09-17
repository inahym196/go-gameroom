package gateway

import (
	"go-gameroom/domain/entity"
)

type RoomRepository struct {
	Rooms []*entity.Room
}

func NewRoomRepository() entity.RoomRepository {
	return &RoomRepository{}
}

func (repo *RoomRepository) GetRoomById(entity.RoomId) (*entity.Room, error) {
	return &entity.Room{}, nil
}

func (repo *RoomRepository) GetRooms() ([]*entity.Room, error) {
	return repo.Rooms, nil
}

func (repo *RoomRepository) Create(entity.RoomId) (*entity.Room, error) {
	return &entity.Room{}, nil
}

func (repo *RoomRepository) Delete(entity.RoomId) error {
	return nil
}
