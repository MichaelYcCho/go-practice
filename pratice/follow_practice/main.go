package main

import (
	"github.com/MichaelYcCho/go-practice/pratice/follow_practice/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitialDB()
	database.Migrate(db)

	defer database.CloseDB()

	// Router setting
	r := gin.Default()
	v1 := r.Group("/api/v1")
	println(v1)
	//v1.Use(middlewares.AuthMiddleware(false))
}
