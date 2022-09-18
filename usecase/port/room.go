package port

type RoomId = int

type Room struct {
	RoomId     RoomId
	Url        string
	Atendee    []string
	GameStatus string
}

type RoomInputPort interface {
	GetRoomById(RoomId) (*Room, error)
	GetRooms() (map[int]*Room, error)
	Init(RoomId) (*Room, error)
}
