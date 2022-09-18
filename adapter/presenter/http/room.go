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
	fmt.Printf("%+v", room)
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, room)
}

func (c *RoomHTTPClient) GetRooms(rooms map[int]*port.Room) {
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, &rooms)
	fmt.Println("===")
	for _, room := range rooms {
		fmt.Printf("%+v\n", room)
	}
}

func (c *RoomHTTPClient) Create(room *port.Room) {
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, &room)
	fmt.Printf("create: %+v", room)
}

func (c *RoomHTTPClient) Delete(ok bool) {
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, ok)
	fmt.Printf("delete %+v", ok)
}
