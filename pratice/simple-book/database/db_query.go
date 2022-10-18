package database

import (
	"github.com/MichaelYcCho/go-practice/pratice/simple-book/models"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return DB
}

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book

	err := db.Debug().Find(&books).Error
	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func DeleteBook(id string, db *gorm.DB) error {
	var b models.Book
	// err := db.Delete(&b).Where(id).Error
	if err := db.Where("id = ?", id).Delete(&b).Error; err != nil {
		// if err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, b *models.Book) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}

func ClearTable() {
	DB.Exec("DELETE FROM books")
	DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}
