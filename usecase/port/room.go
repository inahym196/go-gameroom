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
	Create(RoomId)
	Delete(RoomId)
}

type RoomOutputPort interface {
	GetRoomById(*Room)
	GetRooms(map[int]*Room)
	Create(*Room)
	Delete(bool)
}
