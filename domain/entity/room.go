package entity

type Room struct {
	Id      string
	Url     string
	Atendee []User
	Game    Game
}

type RoomRepository interface {
	GetRoomById(RoomId string) (*Room, error)
	GetRooms() (map[string]*Room, error)
	Init(RoomId string) (*Room, error)
}
