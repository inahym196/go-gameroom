package port

type RoomId = string

type Room struct {
	RoomId     RoomId
	Url        string
	Atendee    []string
	GameStatus string
}

type RoomInputPort interface {
	GetRoomById(RoomId)
	GetRooms()
}

type RoomOutputPort interface {
	GetRoomById(*Room)
	GetRooms([]*Room)
}
