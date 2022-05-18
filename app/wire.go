//go:build wireinject
// +build wireinject

package app

import (
	taskRepo "github.com/evax/app/services/tasks/repository"
	userRepo "github.com/evax/app/services/users/repository"
	userService "github.com/evax/app/services/users/service"
	taskService "github.com/evax/app/services/tasks/service"
	"github.com/google/wire"

	evaxDb "github.com/evax/app/data_sources/evax"
)

func InitialiseApp() *AppImpl {
	panic(wire.Build(
		wire.Bind(new(evaxDb.DataSource), new(*evaxDb.DataSourceImpl)),
		evaxDb.NewEvaxDataSource,

		wire.Bind(new(userRepo.UserRepository), new(*userRepo.UserRepositoryImpl)),
		userRepo.NewUserRepository,

		wire.Bind(new(taskRepo.TasksRepository), new(*taskRepo.TasksRepositoryImpl)),
		taskRepo.NewTaskRepository,

		wire.Bind(new(taskService.TaskService), new(*taskService.TaskServiceImpl)),
		taskService.NewTaskService,

		wire.Bind(new(userService.UserService), new(*userService.UserServiceImpl)),
		userService.NewUserService,

		NewApp,
	))
}
