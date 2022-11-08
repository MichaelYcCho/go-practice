package models

import (
	"github.com/MichaelYcCho/go-practice/pratice/follow_practice/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

type UserModel struct {
	ID           uint          `gorm:"primary_key"`
	Username     string        `gorm:"column:username"`
	Email        string        `gorm:"column:email;unique"`
	Bio          string        `gorm:"column:bio;size:1024"`
	Image        *string       `gorm:"column:image"`
	PasswordHash string        `gorm:"column:password;not null"`
	Followers    []FollowModel `gorm:"foreignkey:FollowingID"`
	Followings   []FollowModel `gorm:"foreignkey:FollowedByID"`
}

// FollowModel M2m Relationship Way
type FollowModel struct {
	gorm.Model
	Following    UserModel
	FollowingID  uint
	FollowedBy   UserModel
	FollowedByID uint
}

func AutoMigrate() {
	db := database.GetDB()

	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&FollowModel{})
}
