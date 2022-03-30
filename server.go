package main

import (
	"github.com/joho/godotenv"
	"github.com/teten-nugraha/go-dev-forum/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// db connection
	config.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
