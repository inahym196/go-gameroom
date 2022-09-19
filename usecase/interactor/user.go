package interactor

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"go-gameroom/usecase/repository"
)

type UserInteractor struct {
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository
}

func NewUserInputPort(userRepository repository.UserRepository, sessionRepository repository.SessionRepository) port.UserInputPort {
	return &UserInteractor{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
	}
}

func (i *UserInteractor) GetUserBySessionId(sessionId string) (*entity.User, error) {
	session, err := i.SessionRepository.Get(sessionId)
	if err != nil {
		return nil, err
	}
	userId := session.Data
	user, err := i.UserRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
