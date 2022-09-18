package entity

type RoomId string

type RoomUrl string

type Room struct {
	Id      RoomId
	Url     RoomUrl
	Atendee []User
	Game    Game
}

type RoomRepository interface {
	GetRoomById(RoomId) (*Room, bool, error)
	GetRooms() (map[int]*Room, error)
	Create(RoomId) (*Room, bool)
	Delete(RoomId) (bool, error)
}
