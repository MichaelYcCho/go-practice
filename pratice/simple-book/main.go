package main

import (
	"log"

	"github.com/MichaelYcCho/go-practice/pratice/simple-book/database"
	"github.com/MichaelYcCho/go-practice/pratice/simple-book/routers"
)

func main() {
	database.ConnectDatabase() // db 연결

	// Proxy 경로 우회를위해 사용
	//gin.SetMode(gin.ReleaseMode)

	r := routers.Setup()
	if err := r.Run("127.0.0.1:4000"); err != nil {
		log.Fatal(err)
	}
}
