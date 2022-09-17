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
	GetRoomById(RoomId) (*Room, error)
	GetRooms() ([]*Room, error)
	Create(RoomId) (*Room, error)
	Delete(RoomId) error
}
