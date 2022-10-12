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

func TestBook_FindBook(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	gormDB := NewGorm(db)

	rows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(testBook.ID, testBook.Title, testBook.Author)

	query := "SELECT * FROM `books` WHERE id = ? ORDER BY `books`.`id` LIMIT 1"
	mock.ExpectQuery(query).WillReturnRows(rows)

	var book Book
	gormDB.Where("id = ?", testBook.ID).First(&book)

	assert.Equal(t, book.Title, testBook.Title)
}

func TestBook_CreateBook(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	gormDB := NewGorm(db)

	query := "INSERT INTO `books` (`title`,`author`) VALUES (?,?,?,?,?)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(testBook.Title, testBook.Author).WillReturnResult(sqlmock.NewResult(1, 1))

	gormDB.Create(&testBook)

	assert.NotEmpty(t, testBook.ID)
	assert.Equal(t, testBook.Title, testBook.Title)
	assert.Equal(t, testBook.Author, testBook.Author)
}
