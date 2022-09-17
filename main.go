package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/roberto-carlos-tg/go-auht/database"
	"github.com/roberto-carlos-tg/go-auht/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":9090")
}
