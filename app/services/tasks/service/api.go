package service

import (
	model "github.com/evax/app/services/tasks/repository"
)

type TaskService interface {
	CreateTask(text string, userID string, userName string) (model.Task, error)
	GetAll() ([]model.Task, error)
}

type TaskServiceImpl struct {
	Repo model.TasksRepository
}

func NewTaskService(userRepo model.TasksRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		Repo: userRepo,
	}
}
