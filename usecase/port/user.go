package port

import (
	"go-gameroom/domain/entity"
)

type UserInputPort interface {
	// userRepository
	// GetUser(UserId) (*entity.User, error)
	// sessionRepository
	// GetData(SessionId) (string, error)
	GetUserBySessionId(sessionId string) (*entity.User, error)
	CreateSession(userName string) (sessionId string, err error)
	DeleteSession(sessionId string) error
	// Auth()
	// Regist()
}
