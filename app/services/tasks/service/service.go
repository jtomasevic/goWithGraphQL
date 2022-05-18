package service

import model "github.com/evax/app/services/tasks/repository"

func (s *TaskServiceImpl) CreateTask(text string, userID string, userName string) (model.Task, error) {
	return s.Repo.Save(model.Task{
		Text: text,
		Done: false,
		UserID: userID,
		Username: userName,
	})
}

func (s *TaskServiceImpl) GetAll() ([]model.Task, error) {
	return s.Repo.All()
}
