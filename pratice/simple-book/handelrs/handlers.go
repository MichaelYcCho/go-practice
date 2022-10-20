package handelrs

import (
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
