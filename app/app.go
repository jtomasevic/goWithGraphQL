package app

import (
	tasks "github.com/evax/app/services/tasks/service"
	userService "github.com/evax/app/services/users/service"
)

type App interface {
	Start()
}

type AppImpl struct {
	UserService userService.UserService
	TaskService tasks.TaskService
}

func (app *AppImpl) Start() {

}

func NewApp(
	userService userService.UserService,
	taskService tasks.TaskService) *AppImpl {
	return &AppImpl{
		UserService: userService,
		TaskService: taskService,
	}
}

var Application *AppImpl

func InitApp() {
	Application = InitialiseApp()
}
