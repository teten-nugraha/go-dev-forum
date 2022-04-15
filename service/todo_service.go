package service

import (
	"github.com/teten-nugraha/go-dev-forum/dto"
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/repository"
)

func LoadTodos() []dto.FetchTodoDto {

	var todoDos []dto.FetchTodoDto

	todos := repository.FetchTodos()

	for i, _ := range todos {
		var dto dto.FetchTodoDto
		dto.Id = todos[i].Id
		dto.Name = todos[i].Nama
		dto.CreatedAt = todos[i].CreatedAt

		todoDos = append(todoDos, dto)
	}

	return todoDos

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

func GetDuration(user_id uint) []models.TodosDone {

	var todos_done []models.TodosDone

	todos := repository.FindAllByUserId(user_id)

	if len(todos) > 0 {
		for i, _ := range todos {
			if todos[i].IsFinish == true {
				var todo_done models.TodosDone

				todo_done.Id = todos[i].Id
				todo_done.Nama = todos[i].Nama
				todo_done.CreatedAt = todos[i].CreatedAt
				todo_done.FinishedAt = todos[i].FinishedAt

				durationTime := todo_done.FinishedAt.Sub(todo_done.CreatedAt)
				duration := int(durationTime.Hours() / 24) // dalam hari
				todo_done.Duration = duration

				todos_done = append(todos_done, todo_done)
			}
		}
	}

	return todos_done
}
