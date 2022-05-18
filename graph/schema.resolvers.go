package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/evax/app"
	"github.com/evax/app/services/auth"
	jwt "github.com/evax/app/services/auth/jwt"
	u "github.com/evax/app/services/users/repository"

	"github.com/evax/graph/generated"
	"github.com/evax/graph/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {

	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Task{}, fmt.Errorf("access denied")
	}

	fmt.Printf("INFO: user id: %s", user.ID)

	task, _ := app.Application.TaskService.CreateTask(input.Text, user.ID, user.Username)
	repoUser, _ := app.Application.UserService.GetUserByID(user.ID)

	returnValue := model.Task{
		Text: task.Text,
		Done: task.Done,
		User: &model.User{Name: repoUser.Username, ID: repoUser.ID},
	}
	return &returnValue, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user u.User
	user.Username = input.Username
	user.Password = input.Password
	app.Application.UserService.CreateUser(user)
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return app.Application.UserService.Authenticate(input.Username, input.Password)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {

	tasks, _ := app.Application.TaskService.GetAll()
	var output []*model.Task
	for _, t := range tasks {
		task := &model.Task{
			ID:   t.ID,
			Text: t.Text,
			Done: t.Done,
			User: &model.User{
				ID:   t.UserID,
				Name: t.Username,
			},
		}
		output = append(output, task)
	}
	return output, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	if app.Application == nil {
		fmt.Println("ERROR: Application is null")
		return []*model.User{}, nil
	}
	if app.Application.UserService == nil {
		fmt.Println("ERROR: Application.UserService is null")
		return []*model.User{}, nil
	}
	usersFromRepo, _ := app.Application.UserService.All()

	var users []*model.User

	for _, u := range usersFromRepo {
		mapedUser := &model.User{
			ID:   u.ID,
			Name: u.Username,
		}
		users = append(users, mapedUser)
	}
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
