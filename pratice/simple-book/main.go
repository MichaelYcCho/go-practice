package main

import (
	"github.com/MichaelYcCho/go-practice/pratice/todo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Proxy 경로 우회를위해 사용
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	models.ConnectDatabase() // new

	r.Run()
}
