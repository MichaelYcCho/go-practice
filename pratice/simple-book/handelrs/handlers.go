package handelrs

import (
	"fmt"
	"github.com/MichaelYcCho/go-practice/pratice/simple-book/models"
	"net/http"

	"github.com/MichaelYcCho/go-practice/pratice/simple-book/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetBooks(c *gin.Context) {
	books, err := database.GetBooks(a.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(books) < 1 {
		c.JSON(http.StatusNotFound, "list is empty")
		return
	}

	c.JSON(http.StatusOK, books)
}

func (a *APIEnv) GetBook(c *gin.Context) {
	id := c.Params.ByName("id")
	book, exists, err := database.GetBookByID(id, a.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no book in db")
		return
	}

	c.JSON(http.StatusOK, book)
}

func (a *APIEnv) CreateBook(c *gin.Context) {
	book := models.Book{}
	err := c.BindJSON(&book)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := a.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}
