package controller

import (
	"fmt"
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"net/http"
	"strconv"
	"strings"
)

type RoomController struct {
	RepositoryFactory func() entity.RoomRepository
	InputFactory      func(repository entity.RoomRepository) port.RoomInputPort
}

func (c *RoomController) EndpointHandler(w http.ResponseWriter, r *http.Request) {
	repository := c.RepositoryFactory()
	inputPort := c.InputFactory(repository)
	roomId := strings.TrimPrefix(r.URL.Path, "/rooms/")
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		if len(roomId) > 0 {
			i, err := strconv.Atoi(roomId)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			inputPort.GetRoomById(i)
			return
		}
		inputPort.GetRooms()
	case http.MethodPost:
		fmt.Println("DELETE")
		i, err := strconv.Atoi(roomId)
		if err != nil {
			fmt.Fprint(w, err)
		}
		inputPort.Init(i)
	}
}
