package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	ID string `gorm:"primaryKey"`
}

type RepostitoryImpl struct {
	db gorm.DB
}

func (repo *RepostitoryImpl) Save(entity Entity) (Entity, error) {
	if entity.ID == "" {
		entity.ID = uuid.NewString()
		repo.db.Create(entity)
	} else {
		repo.db.Save(entity)
	}
	return entity, nil
}

func (repo *RepostitoryImpl) GetByID(id string) (Entity, error) {
	var entity Entity
	repo.db.First(&entity, "ID = ?", id)
	return entity, nil
}

func (repo *RepostitoryImpl) Delete(id string) error {
	var entity Entity
	repo.db.Delete(&entity, "ID = ?", id)
	return nil
}

func (repo *RepostitoryImpl) All() ([]Entity, error) {
	var entities []Entity
	repo.db.Find(&entities)
	return entities, nil
}
