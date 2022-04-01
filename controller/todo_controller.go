package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/service"
	"github.com/teten-nugraha/go-dev-forum/utils"
)

func FetchTodos(c *fiber.Ctx) error {

	todos := service.LoadTodos()

	c.Status(200)
	return c.JSON(fiber.Map{
		"success": true,
		"todos":   todos,
	})

}

func CreateTodo(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")
	idS, _ := utils.ParseJwt(cookie)
	userId, _ := strconv.Atoi(idS)

	newTodo := models.Todo{
		Nama:     data["nama"],
		UserId:   uint(userId),
		IsFinish: false,
	}

	todo := service.CreateTodo(newTodo)
	if (models.Todo{}) == todo {
		c.Status(500)
		return c.JSON(fiber.Map{
			"error":   true,
			"code":    400,
			"message": "Terjadi kesalahan pada saat menyimpan todo",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"success": true,
		"todo":    todo,
	})

}

func UpdateTodo(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	todo := models.Todo{
		Id: uint(id),
	}

	if err := c.BodyParser(&todo); err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"error":   true,
			"code":    500,
			"message": "Terjadi kesalahan pada saat mengolah json parser",
		})
	}

	todoUpdated := service.CreateTodo(todo)
	if (models.Todo{}) == todoUpdated {
		c.Status(500)
		return c.JSON(fiber.Map{
			"error":   true,
			"code":    400,
			"message": "Terjadi kesalahan pada saat menyimpan todo",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"success": true,
		"todo":    todoUpdated,
	})

}

func FinishTodo(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	todoFinished := service.SetFinishTodo(uint(id))

	c.Status(200)
	return c.JSON(fiber.Map{
		"success": true,
		"todo":    todoFinished,
	})
}
