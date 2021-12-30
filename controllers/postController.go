package controllers

import (
	"github.com/denizyoldas/blog-app-api/database"
	"github.com/denizyoldas/blog-app-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	// database.DB.Find(&posts)

	// database.DB.Model(&models.Post{}).
	//Select("users.id, users.name, users.surname, posts.id, posts.title, posts.body").
	//Joins("left join users on posts.user_id = users.id").
	//Scan(&posts)
	database.DB.
		Model(&models.Post{}).
		// Select("posts.id, posts.title, posts.body, users.name").
		Joins("LEFT JOIN users on posts.user_id = users.id").
		Scan(&posts)

	// database.DB.Joins("users").Find(&posts)

	return c.JSON(posts)
}
