package main

import (
	"github.com/MichaelYcCho/go-practice/pratice/todo/controllers"
	"github.com/MichaelYcCho/go-practice/pratice/todo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase() // db 연결

	// Proxy 경로 우회를위해 사용
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	println("Start")

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run(":4000")
}