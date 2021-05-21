package main

import (
	"fmt"

	"github.com/michael_cho77/go-practice/go-scrapper/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("Second")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}
