package presenter

import (
	"fmt"
	"go-gameroom/usecase/port"
	"net/http"
)

type XOGameClient struct {
	writer http.ResponseWriter
}

func NewXOGameOutputPort(w http.ResponseWriter) port.XOGameOutputPort {
	return &XOGameClient{
		writer: w,
	}
}

func (c *XOGameClient) GetGame(game *port.XOGame) {
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, game)
	fmt.Printf("%+v", game)
}
