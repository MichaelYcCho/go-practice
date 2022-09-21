package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	name_list := []string{"John", "Paul", "George", "Ringo"}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", name_list)
	if err != nil {
		log.Fatalln(err)
	}
}
