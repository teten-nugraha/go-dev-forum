package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-dev-forum/controller"
	"github.com/joho/godotenv"
	"log"
	"os"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(server *fiber.App) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SECRET := os.Getenv("SECRET")

	api := server.Group("/api/v1")

	api.Post("/signUp", controller.Register)
	api.Post("/signIn", controller.Login)

	authRoute := api.Group("")
	authRoute.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(SECRET),
	}))

	// profile
	authRoute.Get("/profile", controller.GetProfile)

	// todos
	authRoute.Get("/todos", controller.FetchTodos)
	authRoute.Post("/todos", controller.CreateTodo)
	authRoute.Put("/todos/:id", controller.UpdateTodo)
	authRoute.Post("/todos/finish/:id", controller.FinishTodo)
	authRoute.Get("/todos/done", controller.GetTodoDone)

	// counter
	authRoute.Get("/counters", controller.Counter)
}
