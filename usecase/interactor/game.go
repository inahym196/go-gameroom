package interactor

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
)

type GameInteractor struct {
	OutputPort     port.GameOutputPort
	GameRepository entity.GameRepository
}

func NewGameInputPort(outputPort port.GameOutputPort, gameRepository entity.GameRepository) port.GameInputPort {
	return &GameInteractor{
		OutputPort:     outputPort,
		GameRepository: gameRepository,
	}
}

func (i *GameInteractor) GetGames() {
	res, err := i.GameRepository.FindAll()
	if err != nil {
		panic(0)
	}
	i.OutputPort.GetGames(res)
}
func (i *GameInteractor) GetGame(params *port.GetGameRequestParams) {
	res, err := i.GameRepository.Find(entity.GameId(params.GameId))
	if err != nil {
		panic(0)
	}
	i.OutputPort.GetGame(res)
}
