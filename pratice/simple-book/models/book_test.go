package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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

var testBook = Book{
	ID:     1,
	Title:  "test title",
	Author: "test author",
}

func TestBook_FindBooks(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	gormDB := NewGorm(db)

	rows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(testBook.ID, testBook.Title, testBook.Author)

	query := "SELECT * FROM `books`"
	mock.ExpectQuery(query).WillReturnRows(rows)

	var books []Book
	gormDB.Find(&books)

	assert.Equal(t, books[0].Title, testBook.Title)
}
