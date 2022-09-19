package gateway

import (
	"fmt"
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/repository"
)

type UserRepository struct{}

var UserDataBase map[string]*entity.User = make(map[string]*entity.User)

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) Get(userId string) (*entity.User, error) {
	if user, ok := UserDataBase[userId]; ok {
		fmt.Printf("%#v", user)
		return user, nil
	}
	return nil, fmt.Errorf("User NotFound")
}
