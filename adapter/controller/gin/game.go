package controller

import (
	"go-gameroom/usecase/port"
	"go-gameroom/usecase/repository"
)

type GameController struct {
	RepositoryFactory func() repository.XOGameRepository
	InputPortFactory  func(repository repository.XOGameRepository) port.XOGameInputPort
}

func (c *GameController) GetGames() ([]*port.GetGameResponse, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	games, err := inputport.GetGames()
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (c *GameController) GetGame(id int) (*port.GetGameResponse, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	game, err := inputport.GetGame(id)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (c *GameController) PutPiece(x int, y int, piece string, id int) (*port.PutPieceResponse, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	putpoint := port.NewPoint(x, y)
	res, err := inputport.PutPiece(putpoint, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
