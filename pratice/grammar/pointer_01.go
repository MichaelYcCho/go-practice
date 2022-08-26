package grammar

import "fmt"

func Pointer01() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	changePointer(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func changePointer(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
