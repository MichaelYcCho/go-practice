package controllers

import (
	"net/http"

	"github.com/MichaelYcCho/go-practice/pratice/simple-book/models"
	"github.com/gin-gonic/gin"
)

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})

}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusCreated, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(models.Book{Title: input.Title, Author: input.Author})
	// models.DB.Model(&book).Updates(map[string]interface{}{"Title": input.Title, "Author": input.Author})

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
