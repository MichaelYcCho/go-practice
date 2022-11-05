package handelrs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/MichaelYcCho/go-practice/pratice/simple-book/database"
	"github.com/MichaelYcCho/go-practice/pratice/simple-book/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func insertTestBook(db *gorm.DB) (models.Book, error) {
	b := models.Book{
		Author:    "test",
		Name:      "test",
		PageCount: 10,
	}

	if err := db.Create(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

func Test_GetBooks_Ok(t *testing.T) {
	database.ConnectDatabase()
	db := database.GetDB()
	_, err := insertTestBook(db)
	if err != nil {
		return
	}

	req, w := setGetBooksRouter(db)

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Book{}

	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.Book{}
	a.Equal(expected, actual)
	database.ClearTable()
}

func Test_GetBooks_EmptyResult(t *testing.T) {
	database.ConnectDatabase()
	db := database.GetDB()
	req, w := setGetBooksRouter(db)

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusNotFound, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Book{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.Book{}
	a.Equal(expected, actual)
	database.ClearTable()
}

func Test_GetBook_OK(t *testing.T) {
	a := assert.New(t)
	database.ConnectDatabase()
	db := database.GetDB()

	book, err := insertTestBook(db)
	if err != nil {
		a.Error(err)
	}

	req, w := setGetBookRouter(db, "/"+strconv.Itoa(int(book.ID)))
	println("BookID : ", book.ID)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Book{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := book
	expected.Model = gorm.Model{}
	a.Equal(expected, actual)
	database.ClearTable()
}

func Test_GetBook_InvalidId(t *testing.T) {
	a := assert.New(t)
	database.ConnectDatabase()
	db := database.GetDB()

	req, w := setGetBookRouter(db, "/aa")

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusNotFound, w.Code, "HTTP request status code error")
}

func Test_CreateBook_Ok(t *testing.T) {
	a := assert.New(t)
	database.ConnectDatabase()
	db := database.GetDB()

	book := models.Book{
		Author:    "test",
		Name:      "test",
		PageCount: 20,
	}

	reqBody, err := json.Marshal(book)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateBookRouter(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Book{}
	// 첫번째 인자로 바이트 슬라이스를 넘겨주고, 두번째로 결과를 담게될 변수 포인터를 넘겨준다
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := book
	a.Equal(expected, actual)
	database.ClearTable()

}

func Test_CreateBook_ErrorBind(t *testing.T) {
	type BookTest struct {
		ID        int
		Author    int
		Name      string
		PageCount int
	}
	bookTest := BookTest{
		Author:    3,
		Name:      "test",
		PageCount: 20,
	}

	a := assert.New(t)
	database.ConnectDatabase()
	db := database.GetDB()
	requestBody, err := json.Marshal(bookTest)
	if err != nil {
		a.Error(err)
	}

	request, w, err := setCreateBookRouter(db, bytes.NewBuffer(requestBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, request.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, w.Code, "HTTP request status code error")

}

func Test_CreateBook_ErrorCreate(t *testing.T) {
	// ID 중복으로 인한 생성 실패
	type BookTestCreate struct {
		// gorm.Model
		ID        int
		Author    string
		Name      string
		PageCount int
	}
	bookTestCreate := BookTestCreate{
		ID:        1,
		Author:    "test",
		Name:      "test",
		PageCount: 10,
	}

	a := assert.New(t)
	database.ConnectDatabase()
	db := database.GetDB()

	_, err := insertTestBook(db)
	if err != nil {
		return
	}

	reqBody, err := json.Marshal(bookTestCreate)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateBookRouter(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusInternalServerError, w.Code, "HTTP request status code error")
}

func Test_UpdateBook_IdNotFound(t *testing.T) {
	a := assert.New(t)
	database.ConnectDatabase()
	db := database.GetDB()

	bookTestUpdate := models.Book{
		Author:    "dsdsds",
		Name:      "dsds",
		PageCount: 10,
	}

	bookTestUpdate.ID = 2

	reqBody, err := json.Marshal(bookTestUpdate)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setUpdateBookRouter(db, "/2", bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPut, req.Method, "HTTP request method error")
	a.Equal(http.StatusNotFound, w.Code, "HTTP request status code error")

	database.ClearTable()
}
