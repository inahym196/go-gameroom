package controller

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"net/http"
	"strings"
)

type GameController struct {
	OutputFactory     func(w http.ResponseWriter) port.GameOutputPort
	RepositoryFactory func() entity.GameRepository
	InputFactory      func(outputPort port.GameOutputPort, gameRepository entity.GameRepository) port.GameInputPort
}

func (c *GameController) GetGames(w http.ResponseWriter, r *http.Request) {
	outputPort := c.OutputFactory(w)
	repository := c.RepositoryFactory()
	inputPort := c.InputFactory(outputPort, repository)
	inputPort.GetGames()
}

func (c *GameController) GetGameHandler(w http.ResponseWriter, r *http.Request) {
	outputPort := c.OutputFactory(w)
	repository := c.RepositoryFactory()
	inputPort := c.InputFactory(outputPort, repository)
	gameId := strings.TrimPrefix(r.URL.Path, "/game/")
	params := port.GetGameRequestParams{GameId: gameId}
	inputPort.GetGame(&params)
}
