package models

import (
	"database/sql"
	"fmt"
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

	type mockBehavior func(mock sqlmock.Sqlmock, returnedId int, book Book)

	tests := []struct {
		name         string
		inputBook    Book
		returnedId   int
		mockBehavior mockBehavior
		expectError  bool
	}{
		{
			name: "ok",
			inputBook: Book{
				Title:  "test title",
				Author: "test author",
			},
			returnedId: 1,
			mockBehavior: func(mock sqlmock.Sqlmock, returnedId int, book Book) {
				mock.ExpectQuery("INSERT INTO `books`").
					WithArgs(book.Title, book.Author).
					WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow(returnedId))
			},
			expectError: false,
		},
		{
			name: "Empty field",
			inputBook: Book{
				Title:  "",
				Author: "test author2",
			},
			returnedId: 1,
			mockBehavior: func(mock sqlmock.Sqlmock, returnedId int, book Book) {
				mock.ExpectQuery("INSERT INTO `books`").
					WithArgs(book.Title, book.Author).
					WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow(returnedId).
						RowError(0, fmt.Errorf("error")))
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehavior(mock, test.returnedId, test.inputBook)
			err := gormDB.Create(&test.inputBook).Error
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBook_UpdateBook(t *testing.T) {
	TestBook_CreateBook(t)

	db, mock := NewMock()
	defer db.Close()
	gormDB := NewGorm(db)

	query := "UPDATE `books` SET `title` = ?,`author` = ? WHERE id = ?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(testBook.Title, testBook.Author, testBook.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	gormDB.Save(&testBook)

	println(testBook.ID)
	fmt.Println(testBook.Title)

	assert.NotEmpty(t, testBook.ID)
	assert.Equal(t, testBook.Title, testBook.Title)
	assert.Equal(t, testBook.Author, testBook.Author)

}
