package port

import (
	"go-gameroom/domain/entity"
)

type UserInputPort interface {
	// userRepository
	// GetUser(UserId) (*entity.User, error)
	// sessionRepository
	// GetData(SessionId) (string, error)
	GetUserBySessionId(SessionId string) (*entity.User, error)

	// Auth()
	// Regist()
}
