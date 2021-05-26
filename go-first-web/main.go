package main

import "fmt"

func main() {
	var myString string
	myString = "Green"

	fmt.Println("Origin String is", myString)

	changeUsingPointer(&myString)
	fmt.Println("After String is", myString)
}

func changeUsingPointer(s *string) {
	fmt.Println(s)
	newValue := "Red"
	*s = newValue
}
