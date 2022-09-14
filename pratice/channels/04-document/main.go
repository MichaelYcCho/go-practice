package main

import (
	"fmt"

	"github.com/MichaelYcCho/go-practice/pratice/channels/04-document/mymath"
)

func main() {
	fmt.Println("2 + 3=", mymath.Sum(2, 3))
	fmt.Println("2 + 4=", mymath.Sum(2, 4))
	fmt.Println("2 + 5=", mymath.Sum(2, 5))

}

// godoc -http=:8080
