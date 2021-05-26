package main

import (
	"fmt"
	"log"
)

func main() {
	var whatToSay string
	var saySomethingElse string
	var i int

	whatToSay = saySomething("My World")
	fmt.Println(whatToSay)

	saySomethingElse = saySomething("Bye")
	log.Println(saySomethingElse)

	log.Println(i)
}

func saySomething(s string) string {
	return s

}
