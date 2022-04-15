package dto

import "time"

type FetchTodoDto struct {
	Id           	uint
	Name   			string
	CreatedAt       time.Time
}