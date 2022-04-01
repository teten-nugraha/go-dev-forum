package repository

import (
	"time"

	"github.com/teten-nugraha/go-dev-forum/config"
	"github.com/teten-nugraha/go-dev-forum/models"
)

/**
 * Load todos yang belum finish
 * IsFinish = false
 * order by createdAt desc
 */
func FetchTodos() []models.Todo {

	var todos []models.Todo

	config.DB.Where("is_finish = false").Order("created_at DESC").Find(&todos)

	return todos
}

func CreateTodo(todo models.Todo) models.Todo {

	todo.CreatedAt = time.Now()

	config.DB.Create(&todo)

	return todo
}

func UpdateTodo(todo models.Todo) models.Todo {

	todo.UpdatedAt = time.Now()

	config.DB.Model(&todo).Updates(&todo)

	return todo
}

func SetFinish(id uint64) models.Todo {

	var todo models.Todo

	config.DB.Where("id = ?", id).First(&todo)

	todo.FinishAt = time.Now()

	config.DB.Model(&todo).Updates(&todo)

	return todo

}
