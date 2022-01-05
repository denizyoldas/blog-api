package routes

import (
	"github.com/denizyoldas/blog-app-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Get("/users", controllers.GetUsers)
	app.Get("/posts", controllers.GetPosts)
	app.Post("/posts", controllers.CreatePosts)
}
