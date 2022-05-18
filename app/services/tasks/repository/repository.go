package repository

import "github.com/google/uuid"

func (repo *TasksRepositoryImpl) Save(entity Task) (Task, error) {
	if entity.ID == "" {
		entity.ID = uuid.NewString()
		repo.db.Create(entity)
	} else {
		repo.db.Save(entity)
	}
	return entity, nil
}

func (repo *TasksRepositoryImpl) GetByID(id string) (Task, error) {
	var entity Task
	repo.db.First(&entity, "ID = ?", id)
	return entity, nil
}

func (repo *TasksRepositoryImpl) Delete(id string) error {
	var entity Task
	repo.db.Delete(&entity, "ID = ?", id)
	return nil
}

func (repo *TasksRepositoryImpl) All() ([]Task, error) {
	var entities []Task
	repo.db.Find(&entities)
	return entities, nil
}
