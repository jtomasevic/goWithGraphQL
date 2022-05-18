package repository

import (
	"gorm.io/gorm"
)

type Task struct {
	ID       string `gorm:"primaryKey"`
	Text     string `json:"text"`
	Done     bool   `json:"done"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

func EvaxDbAutoMigrate(db gorm.DB) {
	db.AutoMigrate(&Task{})
}
