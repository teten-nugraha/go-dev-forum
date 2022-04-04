package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/teten-nugraha/go-dev-forum/config"
	"github.com/teten-nugraha/go-dev-forum/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// db connection
	config.Connect()

	// allow cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	// add logger
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// setup routes
	routes.SetupRoutes(app)

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
