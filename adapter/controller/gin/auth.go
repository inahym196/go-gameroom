package controller

import (
	"go-gameroom/usecase/port"
	"go-gameroom/usecase/repository"
)

type UserController struct {
	UserRepositoryFactory    func() repository.UserRepository
	SessionRepositoryFactory func() repository.SessionRepository
	InputPortFactory         func(userRepository repository.UserRepository, sessionRepository repository.SessionRepository) port.UserInputPort
}

func (c *UserController) GetUserName(sessionId string) (UserName string, err error) {
	userRepository := c.UserRepositoryFactory()
	sessionRepository := c.SessionRepositoryFactory()
	inputport := c.InputPortFactory(userRepository, sessionRepository)
	user, err := inputport.GetUserBySessionId(sessionId)
	if err != nil {
		return "", err
	}
	return user.GetName(), nil
}

func (c *UserController) CreateSession(userName string) (string, error) {
	userRepository := c.UserRepositoryFactory()
	sessionRepository := c.SessionRepositoryFactory()
	inputport := c.InputPortFactory(userRepository, sessionRepository)
	sessionId, _ := inputport.CreateSession(userName)
	return sessionId, nil
}

func (c *UserController) DeleteSession(userName string) error {
	userRepository := c.UserRepositoryFactory()
	sessionRepository := c.SessionRepositoryFactory()
	inputport := c.InputPortFactory(userRepository, sessionRepository)
	err := inputport.DeleteSession(userName)
	return err
}
