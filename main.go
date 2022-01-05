package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/denizyoldas/blog-app-api/database"
	"github.com/denizyoldas/blog-app-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/denizyoldas/blog-app-api/docs"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/swagger/*", swagger.Handler)

	routes.Setup(app)

	_ = app.Listen(":3000")
}
