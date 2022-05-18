package repository


import (
	"github.com/google/uuid"
)

func (repo *UserRepositoryImpl) SaveUser(user User) (User, error){
	if user.ID == "" {
	 	user.ID = uuid.NewString()
		repo.db.Create(user)
	} else {
		repo.db.Save(user)
	}
	return user, nil
}

func (repo *UserRepositoryImpl) GetUserByID(id string) (User, error) {
	var user User
	repo.db.First(&user, "ID = ?", id)
	return user, nil
}

func (repo *UserRepositoryImpl) Delete(id string) error{
	var user User
	repo.db.Delete(&user, "ID = ?", id)
	return nil
}

func (repo *UserRepositoryImpl) All() ([]User, error) {
	var users []User
	repo.db.Find(&users)
	return users, nil
}

func (repo *UserRepositoryImpl) GetUserByUsername(username string) (User, error) {
	var user User
	repo.db.First(&user, "Username = ?", username)
	return user, nil
}


