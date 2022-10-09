package models

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func NewGorm(db *sql.DB) *gorm.DB {
	dbDirver := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dbDirver, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gormDB
}
