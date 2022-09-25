package controllers

import (
	"net/http"

	"github.com/MichaelYcCho/go-practice/pratice/todo/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
