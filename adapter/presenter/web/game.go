package presenter

import (
	"fmt"
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"net/http"
)

type WebClient struct {
	writer http.ResponseWriter
}

func NewGameOutputPort(w http.ResponseWriter) port.GameOutputPort {
	return &WebClient{
		writer: w,
	}
}

func (c *WebClient) GetGames(games []entity.Game) {
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, games)
}

func (c *WebClient) GetGame(game *entity.Game) {
	c.writer.WriteHeader(http.StatusOK)
	fmt.Fprint(c.writer, game)
	fmt.Printf("%+v", game)
}
