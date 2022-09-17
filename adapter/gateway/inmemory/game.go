package gateway

import (
	"fmt"
	"go-gameroom/domain/entity"
)

type XOGameRepository struct {
	XOGameData entity.XOGame
}

func NewXOGameRepository() entity.XOGameRepository {
	return &XOGameRepository{}
}

func (repo *XOGameRepository) Init() error {
	fmt.Printf("%+v", repo)
	return nil
}

func (repo *XOGameRepository) Join(u *entity.User) bool {
	fmt.Printf("%+v", repo)
	return true
}

func (repo *XOGameRepository) Get() (*entity.XOGame, error) {
	return &entity.XOGame{}, nil
}

func (repo *XOGameRepository) PutPiece(p *entity.PutPoint) (*entity.XOGame, error) {
	return &entity.XOGame{}, nil
}
