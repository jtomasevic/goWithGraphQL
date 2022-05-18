package repository

import (
	data "github.com/evax/app/data_sources/evax"
	"gorm.io/gorm"
)

type TasksRepository interface {
	Save(entity Task) (Task, error)
	GetByID(id string) (Task, error)
	Delete(id string) error
	All() ([]Task, error)
}

type TasksRepositoryImpl struct {
	db   gorm.DB
	base RepostitoryImpl
}

func NewTaskRepository(dataSource data.DataSource) *TasksRepositoryImpl {
	return &TasksRepositoryImpl{
		db: *dataSource.DB(),
		base: RepostitoryImpl{
			db: *dataSource.DB(),
		},
	}
}
