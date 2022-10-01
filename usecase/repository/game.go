package repository

import "go-gameroom/domain/entity"

type GameRepository interface {
	Init() error
	Join(*entity.User) bool
}

type XOGameRepository interface {
	GameRepository
	GetGame(id int) (*entity.XOGame, error)
	GetGames() ([]*entity.XOGame, error)
	PutPiece(x, y, id int) (*entity.XOGame, error)
}
