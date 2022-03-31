package service

import (
	"errors"
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/repository"
	"github.com/teten-nugraha/go-dev-forum/utils"
	"strconv"
)

func SignUp(username, email, password, password_confirmation string) (models.User, error) {
	var userExist = repository.FindUserByEmail(email)

	// check user exist
	if (models.User{}) != userExist {
		return models.User{}, errors.New("User dengan email " + email + " sudah terdaftar")
	}

	// check password and password confirmation
	if password != password_confirmation {
		return models.User{}, errors.New("Password tidak cocok, silahkan diubah dulu.")
	}

	user := models.User{
		Username: username,
		Email:    email,
	}
	user.SetPassword(password)

	repository.CreateUser(user)

	return user, nil
}

func SignIn(email, password string) (string, error) {

	// check user exist
	var userExist = repository.FindUserByEmail(email)
	if (models.User{}) == userExist {
		return "", errors.New("User dengan email " + email + " tidak terdaftar")
	}

	// compare password
	if err := userExist.ComparePassword(password); err != nil {
		return "", errors.New("Password tidak cocok silahkan diperbaiki dulu.")
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(userExist.Id)))
	if err != nil {
		return "", errors.New("Terjadi kesalahan pada saat generate token")
	}

	return token, nil

}
