package database

import (
	"github.com/MichaelYcCho/go-practice/pratice/follow_practice/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	models.AutoMigrate()

}
