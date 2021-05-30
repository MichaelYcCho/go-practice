package main

import (
	"fmt"

	"github.com/michael_cho77/go-practice/go-first-web/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	RandomNumber := helpers.RandomNumber(numPool)
	intChan <- RandomNumber

}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	for i := 0; i < 2; i++ {
		go CalculateValue(intChan)
	}
	for j := 0; j < 2; j++ {
		fmt.Println(<-intChan)
	}
}
