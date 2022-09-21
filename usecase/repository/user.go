package repository

import (
	"go-gameroom/domain/entity"
)

type UserRepository interface {
	Get(Id string) (*entity.User, error)
	// Auth(*entity.User) (bool, error)
	// Regist(*entity.User) error
}
