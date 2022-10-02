package gateway

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/repository"
)

type XOGameRepository struct{}

var GameDataBase []*entity.XOGame = []*entity.XOGame{entity.NewXOGame(0), entity.NewXOGame(1), entity.NewXOGame(2), entity.NewXOGame(3)}

func NewXOGameRepository() repository.XOGameRepository {
	return &XOGameRepository{}
}

func (repo *XOGameRepository) Init() error {
	return nil
}

func (repo *XOGameRepository) Join(u *entity.User) bool {
	return true
}

func (repo *XOGameRepository) GetGame(id int) (*entity.XOGame, error) {
	return GameDataBase[id], nil
}

func (repo *XOGameRepository) GetGames() ([]*entity.XOGame, error) {
	return GameDataBase, nil
}

func (repo *XOGameRepository) InitGame(id int, owner string) (*entity.XOGame, error) {
	GameDataBase[id] = entity.NewXOGame(id)
	GameDataBase[id].BecomeOwner(owner)
	return GameDataBase[id], nil
}

func (repo *XOGameRepository) SelectGameOrder(id int, user string, order string) (*entity.XOGame, error) {
	err := GameDataBase[id].SetPlayer(user, order)
	if err != nil {
		return nil, err
	}
	return GameDataBase[id], nil
}

func (repo *XOGameRepository) JoinGame(id int, user string) (*entity.XOGame, error) {
	first, draw := GameDataBase[id].GetPlayers()
	if first == "" {
		GameDataBase[id].SetPlayer(user, "first")
	} else if draw == "" {
		GameDataBase[id].SetPlayer(user, "draw")
	} else {
		panic(0)
	}
	return GameDataBase[id], nil
}

func (repo *XOGameRepository) PutPiece(id, x, y int, piece string) (*entity.XOGame, error) {
	GameDataBase[id].SetPiece(x, y, piece)
	return GameDataBase[id], nil
}
