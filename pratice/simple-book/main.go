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

	r.Run(":4000")
}
