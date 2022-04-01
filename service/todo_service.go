package service

import (
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/repository"
)

func LoadTodos() []models.Todo {

	todos := repository.FetchTodos()

	return todos

}

func CreateTodo(newTodo models.Todo) models.Todo {

	todo := repository.CreateTodo(newTodo)

	return todo

}

func UpdateTodo(todo models.Todo) models.Todo {

	todoUpdated := repository.UpdateTodo(todo)

	return todoUpdated
}

func SetFinishTodo(id uint) models.Todo {

	todoFinished := repository.SetFinish(id)

	return todoFinished

}
