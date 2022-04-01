package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-dev-forum/controller"
	"github.com/teten-nugraha/go-dev-forum/middlewares"
)

func SetupRoutes(server *fiber.App) {

	api := server.Group("/api/v1")

	api.Post("/signUp", controller.Register)
	api.Post("/signIn", controller.Login)

	authRoute := api.Use(middlewares.IsAuthenticated)

	// profile
	authRoute.Get("/profile", controller.GetProfile)

	// todos
	authRoute.Get("/todos", controller.FetchTodos)
	authRoute.Post("/todos", controller.CreateTodo)
	authRoute.Put("/todos/:id", controller.UpdateTodo)
	authRoute.Post("/todos/finish/:id", controller.FinishTodo)

	// counter
	authRoute.Get("/counters", controller.Counter)
}
