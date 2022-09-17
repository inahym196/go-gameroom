package presenter

import (
	"fmt"
	"go-gameroom/usecase/port"
	"net/http"
)

type RoomHTTPClient struct {
	writer http.ResponseWriter
}

func NewRoomOutputPort(w http.ResponseWriter) port.RoomOutputPort {
	return &RoomHTTPClient{
		writer: w,
	}
}

func (c *RoomHTTPClient) GetRoomById(room *port.Room) {
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, room)
	fmt.Printf("%+v", room)
}

func (c *RoomHTTPClient) GetRooms(rooms []*port.Room) {
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, &rooms)
	for _, room := range rooms {
		fmt.Printf("%#v", room)
	}
}
