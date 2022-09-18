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

func (c *RoomHTTPClient) GetRoomById(room *port.Room, err error) {
	c.writer.WriteHeader(http.StatusOK)
	if err != nil {
		fmt.Fprint(c.writer, err)
	}
	fmt.Printf("%+v", room)
	fmt.Fprint(c.writer, room)
}

func (c *RoomHTTPClient) GetRooms(rooms map[int]*port.Room, err error) {
	c.writer.WriteHeader(http.StatusOK)
	if err != nil {
		fmt.Fprint(c.writer, err)
	}
	for _, room := range rooms {
		fmt.Printf("%+v\n", room)
	}
	fmt.Fprint(c.writer, &rooms)
}

func (c *RoomHTTPClient) Init(room *port.Room, err error) {
	c.writer.WriteHeader(http.StatusOK)
	if err != nil {
		fmt.Fprint(c.writer, err)
	}
	fmt.Printf("%+v", room)
	fmt.Fprint(c.writer, &room)
}
