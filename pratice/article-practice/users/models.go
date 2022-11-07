package users

import (
	"errors"
	"github.com/MichaelYcCho/go-practice/pratice/article-practice/config"
	"golang.org/x/crypto/bcrypt"
)

// UserModel is the model of user table
//
//	If you want to split null and "", you should use *string instead of string.
type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

func AutoMigrate() {
	db := config.GetDB()
	db.AutoMigrate(&UserModel{})
}

// About Bcrypt:  https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
func (user *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.PasswordHash = string(passwordHash)
	return nil
}

func (user *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(user.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// userModel, err := FindOneUser(&UserModel{Username: "naming"})
func FindOneUser(condition interface{}) (UserModel, error) {
	db := config.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// if err := SaveModel(&userModel); err != nil { ... }
func SaveModel(data interface{}) error {
	db := config.GetDB()
	err := db.Save(data).Error
	return err
}

// err := db.Model(userModel).Update(UserModel{Username: "naming"}).Error
func (user *UserModel) Update(data interface{}) error {
	db := config.GetDB()
	err := db.Model(user).Updates(data).Error
	return err
}
