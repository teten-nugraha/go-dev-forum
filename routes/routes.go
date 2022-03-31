package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-dev-forum/controller"
)

func SetupRoutes(server *fiber.App) {

	api := server.Group("/api/v1")

	api.Post("/signUp", controller.Register)
	api.Post("/signIn", controller.Login)
}
