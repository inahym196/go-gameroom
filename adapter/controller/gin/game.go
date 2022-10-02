package controller

import (
	"fmt"
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

func (c *GameController) InitGame(id int, owner string) (*port.GetGameResponse, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	game, err := inputport.InitGame(id, owner)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (c *GameController) JoinGame(id int, user string) (*port.GetGameResponse, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	game, err := inputport.GetGame(id)
	if err != nil {
		return nil, err
	}
	switch game.Status {
	case "Init":
		game, err = inputport.InitGame(id, user)
	case "Waiting":
		if game.Owner != user {
			game, err = inputport.JoinGame(id, user)
			fmt.Println("second user joined")
			fmt.Printf("%#v", game)
		}
	}
	return game, nil
}

func (c *GameController) SelectGameOrder(id int, user string, order string) (*port.GetGameResponse, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	game, err := inputport.GetGame(id)
	if !(game.Status == "Setting" && game.Owner == user) {
		return nil, fmt.Errorf("invalid operation")
	}
	game, err = inputport.SelectGameOrder(id, user, order)
	if err != nil {
		return nil, err
	}
	return game, nil
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

func (c *GameController) PutPiece(x int, y int, id int, user string) (*port.PutPieceResponse, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	putpoint := port.NewPoint(x, y)
	res, err := inputport.PutPiece(id, putpoint, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
