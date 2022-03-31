package repository

import (
	"github.com/teten-nugraha/go-dev-forum/config"
	"github.com/teten-nugraha/go-dev-forum/models"
)

func FindUserByEmail(email string) models.User {

	var user models.User

	config.DB.Where("email = ?", email).Find(&user)

	return user
}
