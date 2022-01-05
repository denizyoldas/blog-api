package controllers

import (
	"github.com/denizyoldas/blog-app-api/database"
	"github.com/denizyoldas/blog-app-api/models"
	"github.com/gofiber/fiber/v2"
)

// GetPosts godoc
// @Summary Show a models.Post
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Post
// @Router /posts [get]
func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	database.DB.
		Model(&models.Post{}).
		Preload("User").Find(&posts)

	return c.JSON(posts)
}

// CreatePosts godoc
// @Summary Show a models.Post
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Post
// @Router /posts [post]
func CreatePosts(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	return nil
}
