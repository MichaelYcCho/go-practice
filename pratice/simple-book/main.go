package main

import (
	"github.com/MichaelYcCho/go-practice/pratice/todo/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.ConnectDatabase() // new

	r.Run()
}
