package interactor

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
)

type XOGameInteractor struct {
	OutputPort port.XOGameOutputPort
	Repository entity.XOGameRepository
}

func NewXOGameInputPort(outputPort port.XOGameOutputPort, repository entity.XOGameRepository) port.XOGameInputPort {
	return &XOGameInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func gameTransfer(entityGame *entity.XOGame) (outputGame *port.XOGame) {
	return &port.XOGame{}
}

func (i *XOGameInteractor) GetGame() {
	res, err := i.Repository.Get()
	if err != nil {
		panic(0)
	}
	outputGame := gameTransfer(res)
	i.OutputPort.GetGame(outputGame)
}
