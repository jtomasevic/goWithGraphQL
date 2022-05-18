package repository

import (
	model "github.com/evax/app/services/tasks/repository"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	FirstName *string
	LastName  *string
	Username  string `json:"username"`
	Password  string
	Tasks     []model.Task `gorm:"foreignKey:UserID"`
}

