package service

import (
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/repository"
)

func CalculateCounter(user_id uint) models.Counter {

	todos := repository.FindAllByUserId(user_id)

	counter := models.Counter{
		CountTodo:           len(todos),
		CountTodoFinished:   0,
		Sisa:                0,
		ProductivityPercent: 0.0,
	}

	for i, _ := range todos {
		if todos[i].IsFinish == true {
			counter.CountTodoFinished++
		}
	}

	counter.Sisa = counter.CountTodo - counter.CountTodoFinished
	counter.ProductivityPercent = float32(counter.CountTodo / counter.CountTodoFinished)

	return counter
}
