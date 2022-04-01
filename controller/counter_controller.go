package controller

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-dev-forum/service"
	"github.com/teten-nugraha/go-dev-forum/utils"
)

func Counter(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")
	idS, _ := utils.ParseJwt(cookie)
	userId, _ := strconv.Atoi(idS)

	counter := service.CalculateCounter((uint(userId)))

	c.Status(http.StatusOK)
	return c.JSON(fiber.Map{
		"success": true,
		"counter": counter,
	})
}
