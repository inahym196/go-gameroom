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

func (repo *XOGameRepository) PutPiece(x, y, id int) (*entity.XOGame, error) {
	GameDataBase[id].SetPiece(x, y, "XG")
	return GameDataBase[id], nil
}
