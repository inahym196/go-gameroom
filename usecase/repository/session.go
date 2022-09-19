package repository

import (
	"go-gameroom/domain/entity"
)

type SessionRepository interface {
	Get(id string) (*entity.Session, error)
	Create(data string) (*entity.Session, error)
}
