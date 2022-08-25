package grammar

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Factorial() {
	n := factorial(4)
	fmt.Println(n)
}
