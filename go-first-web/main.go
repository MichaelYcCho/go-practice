package main

import (
	"fmt"
)

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}
	fmt.Println(myVar.printFirstName())
	fmt.Println(myVar2.printFirstName())

}
