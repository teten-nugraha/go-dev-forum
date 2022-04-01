package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-dev-forum/service"
	"github.com/teten-nugraha/go-dev-forum/utils"
)

func GetProfile(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")
	idS, _ := utils.ParseJwt(cookie)
	userId, _ := strconv.Atoi(idS)

	user, err := service.Profile(uint(userId))
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	c.Status(http.StatusOK)
	return c.JSON(fiber.Map{
		"success": true,
		"profile": user,
	})
}
