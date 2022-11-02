package handelrs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/MichaelYcCho/go-practice/pratice/simple-book/database"
	"github.com/MichaelYcCho/go-practice/pratice/simple-book/models"
	"github.com/gin-gonic/gin"
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

func setGetBooksRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	api := &APIEnv{DB: db}
	r.GET("/", api.GetBooks)
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
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

func setGetBookRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	api := &APIEnv{DB: db}
	r.GET("/:id", api.GetBook)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
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

func setCreateBookRouter(db *gorm.DB,
	body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	api := &APIEnv{DB: db}
	r.POST("/", api.CreateBook)
	req, err := http.NewRequest(http.MethodPost, "/", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
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
