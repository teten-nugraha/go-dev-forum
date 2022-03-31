package service

import (
	"errors"
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/repository"
)

func Profile(id int64) (models.User, error) {

	var userExist = repository.FindUserById(id)
	if (models.User{}) == userExist {
		return models.User{}, errors.New("User tidak ada di db")
	}

	return userExist, nil
}
