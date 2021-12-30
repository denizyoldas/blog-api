package main

import (
	"github.com/denizyoldas/blog-app-api/database"
	"github.com/denizyoldas/blog-app-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	_ = app.Listen(":3000")
}
