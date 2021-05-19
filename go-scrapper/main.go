package main

import "fmt"

func main() {
	names := []string{"one", "two", "three"}
	names = append(names, "Four")
	fmt.Println(names)
}
