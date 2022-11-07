package config

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./../go-article.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(3)
	DB = db
	return DB
}

// This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	testDB, err := gorm.Open(sqlite.Open("./../gorm_test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("db err: (TestDBInit) ", err)
	}

	sqlDB, _ := testDB.DB()
	sqlDB.SetMaxIdleConns(3)

	DB = testDB
	return DB
}

// Delete the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	err := os.Remove("./../gorm_test.db")
	return err
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
