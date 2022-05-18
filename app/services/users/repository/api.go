package repository

import (
	data "github.com/evax/app/data_sources/evax"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveUser(user User) (User, error)
	GetUserByID(id string) (User, error)
	Delete(id string) error
	All() ([]User, error)
	GetUserByUsername(username string) (User, error)
}	

type UserRepositoryImpl struct {
	db gorm.DB
}

func NewUserRepository(dataSource data.DataSource) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: *dataSource.DB(),
	}
}

