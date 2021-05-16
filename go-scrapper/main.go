package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return lenght, uppercase

}

func main() {
	defer fmt.Print("main done")
	totalLenght, up := lenAndUpper("michael")
	fmt.Println(totalLenght, up)
}
