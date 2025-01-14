package services

import (
	"github.com/heartlezz7/go_gin_todolist/model"
	reporepository "github.com/heartlezz7/go_gin_todolist/repositories"
)

func RegisterService(user model.CreateUser) error {

	err := reporepository.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func LoginService(input string, data *model.UserData) error {

	err := reporepository.Login(input, data)
	if err != nil {
		return err
	}

	return nil

}
