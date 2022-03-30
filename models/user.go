package models

type User struct {
	Id       int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
