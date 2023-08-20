package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/michael_cho77/go-web-base-practice/pkg/config"
	"github.com/michael_cho77/go-web-base-practice/pkg/handlers"
	"github.com/michael_cho77/go-web-base-practice/pkg/render"
)

const portNumber = ":8000"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
