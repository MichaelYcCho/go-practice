package handelrs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
