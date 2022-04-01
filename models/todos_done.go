package models

import "time"

type TodosDone struct {
	Id         uint      `json:"id"`
	Nama       string    `json:"nama"`
	CreatedAt  time.Time `json:"created_at"`
	FinishedAt time.Time `json:"finished_at"`
	Duration   int       `json:"day_duration"`
}
