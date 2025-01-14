package services

import (
	"github.com/heartlezz7/go_gin_todolist/model"
	reporepository "github.com/heartlezz7/go_gin_todolist/repositories"
	"golang.org/x/crypto/bcrypt"
)

func RegisterService(user model.CreateUser) error {

	passwordHash, hassErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if hassErr != nil {
		return hassErr
	}

	user.Password = string(passwordHash)

	err := reporepository.Register(&user)
	if err != nil {
		return err
	}
	return nil
}

func LoginService(email string, data *model.UserData) error {

	err := reporepository.Login(email, data)
	if err != nil {
		return err
	}

	return nil

}
