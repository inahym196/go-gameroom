package gateway

import (
	"fmt"
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/repository"
)

type UserRepository struct{}

var UserDataBase map[string]*entity.User = map[string]*entity.User{
	"user1": entity.NewUser("user1"),
	"user2": entity.NewUser("user2"),
	"user3": entity.NewUser("user3"),
	"user4": entity.NewUser("user4"),
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Get(userName string) (*entity.User, error) {
	if user, ok := UserDataBase[userName]; ok {
		fmt.Printf("%#v", user)
		return user, nil
	}
	return nil, fmt.Errorf("User NotFound")
}
