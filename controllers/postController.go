package controllers

import (
	"github.com/denizyoldas/blog-app-api/database"
	"github.com/denizyoldas/blog-app-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	database.DB.
		Model(&models.Post{}).
		Preload("User").Find(&posts)

	return c.JSON(posts)
}
