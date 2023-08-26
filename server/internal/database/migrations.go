package database

import (
	"github.com/joomcold/go-next-chat/internal/app/models"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate models")
	}
}
