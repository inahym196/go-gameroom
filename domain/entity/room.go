package entity

import "strconv"

type RoomId string

func (roomId RoomId) MustInt() int {
	i, err := strconv.Atoi(string(roomId))
	if err != nil {
		panic(0)
	}
	return i
}

type RoomUrl string

type Room struct {
	Id      RoomId
	Url     RoomUrl
	Atendee []User
	Game    Game
}

type RoomRepository interface {
	GetRoomById(RoomId) (*Room, error)
	GetRooms() (map[int]*Room, error)
	Init(RoomId) (*Room, error)
}
