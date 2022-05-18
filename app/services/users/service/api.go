package service

import (

	model "github.com/evax/app/services/users/repository"
)

type UserService interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByID(id string) (model.User, error)
	Delete(id string) error
	All() ([]model.User, error)
	GetUserIdByUsername(username string) (string, error)
	Authenticate(userName string, password string) (string, error)
}

type UserServiceImpl struct {
	Repo model.UserRepository
}

func NewUserService(userRepo model.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		Repo: userRepo,
	}
}
