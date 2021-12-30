package database

import (
	"github.com/denizyoldas/blog-app-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	con, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/blogdb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("database connection error")
	}

	DB = con

	_ = con.AutoMigrate(&models.User{}, &models.Post{})
}
