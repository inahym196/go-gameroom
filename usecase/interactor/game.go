package interactor

import (
	"fmt"
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
		Owner:  game.GetOwner(),
		Status: game.GetStatus(),
	}
	return gameResponse, nil
}

func (i *XOGameInteractor) InitGame(id int, owner string) (*port.GetGameResponse, error) {
	game, err := i.Repository.InitGame(id, owner)
	if err != nil {
		panic(0)
	}
	gameResponse := &port.GetGameResponse{
		Id:     game.GetId(),
		Owner:  game.GetOwner(),
		Status: game.GetStatus(),
	}
	fmt.Println(gameResponse)
	return gameResponse, nil
}

func (i *XOGameInteractor) JoinGame(id int, user string) (*port.GetGameResponse, error) {
	game, err := i.Repository.JoinGame(id, user)
	if err != nil {
		panic(0)
	}
	gameResponse := &port.GetGameResponse{
		Id:          game.GetId(),
		Owner:       game.GetOwner(),
		Status:      game.GetStatus(),
		FirstPlayer: game.GetPlayerByOrder("first"),
		DrawPlayer:  game.GetPlayerByOrder("draw"),
		Turn:        game.GetTurn(),
	}
	return gameResponse, nil
}

func (i *XOGameInteractor) SelectGameOrder(id int, user string, order string) (*port.GetGameResponse, error) {
	game, _ := i.Repository.SelectGameOrder(id, user, order)
	gameResponse := &port.GetGameResponse{
		Id:          game.GetId(),
		Owner:       game.GetOwner(),
		Status:      game.GetStatus(),
		FirstPlayer: game.GetPlayerByOrder("first"),
		DrawPlayer:  game.GetPlayerByOrder("draw"),
		Turn:        game.GetTurn(),
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

func (i *XOGameInteractor) PutPiece(id int, p *port.Point, user string) (*port.PutPieceResponse, error) {
	game, err := i.Repository.GetGame(id)
	if err != nil {
		return nil, err
	}
	piece := game.GetNextPiece()
	game, err = i.Repository.PutPiece(id, p.X, p.Y, piece)
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
		Status: game.GetStatus(),
		Turn:   game.GetTurn(),
	}
	return putPieceResponse, nil
}
