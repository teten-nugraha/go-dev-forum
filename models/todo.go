package models

import "time"

type Todo struct {
	Id        int64     `json:"id" gorm:"primary_key;auto_increment;not_null"`
	UserId    int64     `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserId"`
	Nama      string    `json:"nama"`
	IsFinish  bool      `json:"is_finish"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
