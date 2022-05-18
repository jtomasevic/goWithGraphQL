package service

import (
	jwt "github.com/evax/app/services/auth/jwt"

	model "github.com/evax/app/services/users/repository"
	"github.com/evax/app/services/users/service/errors"
	"golang.org/x/crypto/bcrypt"
)

func (service *UserServiceImpl) CreateUser(user model.User) (model.User, error) {
	hashedPassword, _ := HashPassword(user.Password)
	user.Password = hashedPassword
	user, _ = service.Repo.SaveUser(user)
	return user, nil
}

func (service *UserServiceImpl) GetUserByID(id string) (model.User, error) {
	return service.Repo.GetUserByID(id)
}

func (service *UserServiceImpl) Delete(id string) error {
	return service.Repo.Delete(id)
}

func (service *UserServiceImpl) All() ([]model.User, error) {
	return service.Repo.All()
}

func (service *UserServiceImpl) GetUserIdByUsername(username string) (string, error) {
	user, _ := service.Repo.GetUserByUsername(username)
	return user.ID, nil
}

func (service *UserServiceImpl) Authenticate(username string, password string) (string, error) {
	user, _ := service.Repo.GetUserByUsername(username)
	if user.ID == "" {
		return "", &errors.WrongUsernameError{}
	}
	correct := CheckPasswordHash(password, user.Password)
	if correct {
		token, err := jwt.GenerateToken(user.Username)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", &errors.WrongPasswordError{}
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
