package database

import (
	"os"

	"github.com/MichaelYcCho/go-practice/pratice/simple-book/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(127.0.0.1:3306)/gin_book_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db

	modelList := []interface{}{
		&models.Book{},
	}
	// db.AutoMigrate(&Book{})
	// db.AutoMigrate(&User{})
	db.AutoMigrate(modelList...)
	//AutoMigrateWithTransaction(db, modelList)
	// GetTableList()
	// DropUnusedColumns(DB, modelList)

}
