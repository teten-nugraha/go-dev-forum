package service

import (
	"errors"
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/repository"
)

func SignUp(username, email, password, password_confirmation string) (models.User, error) {
	var userExist = repository.FindUserByEmail(email)
	if (models.User{}) != userExist {
		return models.User{}, errors.New("User dengan email " + email + " sudah terdaftar")
	}

	user := models.User{
		Username: "Teten",
		Email:    "teten@mail.com",
		Password: "123",
	}

	return user, nil
}
