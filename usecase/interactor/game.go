package interactor

import (
	"go-gameroom/usecase/port"
	"go-gameroom/usecase/repository"
)

type XOGameInteractor struct {
	Repository repository.XOGameRepository
}

func NewXOGameInputPort(repository repository.XOGameRepository) port.XOGameInputPort {
	return &XOGameInteractor{
		Repository: repository,
	}
}

func (i *XOGameInteractor) GetGame(id int) (*port.GetGameResponse, error) {
	game, err := i.Repository.GetGame(id)
	if err != nil {
		panic(0)
	}
	gameResponse := &port.GetGameResponse{
		Id:     game.GetId(),
		Status: game.GetStatus(),
	}
	return gameResponse, nil
}

func (i *XOGameInteractor) GetGames() ([]*port.GetGameResponse, error) {
	game, err := i.Repository.GetGames()
	if err != nil {
		panic(0)
	}
	var gameResponses []*port.GetGameResponse
	for _, g := range game {
		gameResponse := &port.GetGameResponse{
			Id:     g.GetId(),
			Status: g.GetStatus(),
		}
		gameResponses = append(gameResponses, gameResponse)
	}
	return gameResponses, nil
}

func (i *XOGameInteractor) PutPiece(p *port.Point, id int) (*port.PutPieceResponse, error) {
	game, err := i.Repository.PutPiece(p.X, p.Y, id)
	if err != nil {
		return nil, err
	}
	var piecesStr [10][10]string
	for i, col := range game.GetPieces() {
		for j, raw := range col {
			piecesStr[i][j] = raw.String()
		}
	}
	putPieceResponse := &port.PutPieceResponse{
		Pieces: &piecesStr,
	}
	return putPieceResponse, nil
}
