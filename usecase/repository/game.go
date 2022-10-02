package repository

import "go-gameroom/domain/entity"

type GameRepository interface {
	Init() error
	Join(*entity.User) bool
}

type XOGameRepository interface {
	GameRepository
	GetGame(id int) (*entity.XOGame, error)
	InitGame(id int, owner string) (*entity.XOGame, error)
	JoinGame(id int, user string) (*entity.XOGame, error)
	GetGames() ([]*entity.XOGame, error)
	SelectGameOrder(id int, user string, order string) (*entity.XOGame, error)
	PutPiece(id, x, y int, piece string) (*entity.XOGame, error)
}
