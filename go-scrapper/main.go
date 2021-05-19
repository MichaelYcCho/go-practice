package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 30
	fmt.Println(a, *b)
}
