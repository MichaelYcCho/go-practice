package dblayer

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_id := os.Getenv("DB_ID")
	db_name := os.Getenv("DB_NAME")
	db_pwd := os.Getenv("DB_PASSWORD")

	dsn := db_id + ":" + db_pwd + "@tcp(127.0.0.1:3306)/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &DBORM{
		DB: db,
	}, err
}
