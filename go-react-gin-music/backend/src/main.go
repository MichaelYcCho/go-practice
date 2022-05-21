package main

import (
	"log"

	"github.com/MichaelYcCho/go-practice/go-react-gin-music/backend/src/rest"
)

func main() {
	log.Println("Main log....")
	rest.RunAPI("127.0.0.1:8000")
}
