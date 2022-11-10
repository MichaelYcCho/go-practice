package database

import (
	"fmt"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitialDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(env.ConnectionString), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(env.ConnectionString)
		panic("Can't connect to DB!")
	}
	//db.LogMode(true)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't close to DB!")
	}
	sqlDB.Close()
}
