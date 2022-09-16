package presenter

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
)

type GamePresenter struct {
	OutputPort port.GameOutputPort
}

func NewGamePresenter() *GamePresenter {
	return &GamePresenter{}
}

func (p *GamePresenter) GetGames(games []entity.Game) (*port.GetGamesResponse, error) {
	return &port.GetGamesResponse{games}, nil
}

func (p *GamePresenter) GetGame(game *entity.Game) (*port.GetGameResponse, error) {
	return &port.GetGameResponse{game}, nil
}
