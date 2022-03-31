package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-dev-forum/models"
	"github.com/teten-nugraha/go-dev-forum/service"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	user, err := service.SignUp(data["username"], data["email"], data["password"], data["password_confirmation"])
	if err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"error":   true,
			"code":    401,
			"message": err.Error(),
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
