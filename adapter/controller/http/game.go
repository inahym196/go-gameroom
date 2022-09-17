package controller

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"net/http"
)

type GameController struct {
	OutputFactory     func(w http.ResponseWriter) port.XOGameOutputPort
	RepositoryFactory func() entity.XOGameRepository
	InputFactory      func(outputPort port.XOGameOutputPort, gameRepository entity.XOGameRepository) port.XOGameInputPort
}

func (c *GameController) GetGameHandler(w http.ResponseWriter, r *http.Request) {
	outputPort := c.OutputFactory(w)
	repository := c.RepositoryFactory()
	inputPort := c.InputFactory(outputPort, repository)
	inputPort.GetGame()
}
