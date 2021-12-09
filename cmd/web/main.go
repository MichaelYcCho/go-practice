package main

import (
	"fmt"
	"net/http"

	"github.com/yc/go-web-base-practice/pkg/handlers"
)

const portNumber = ":8000"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
