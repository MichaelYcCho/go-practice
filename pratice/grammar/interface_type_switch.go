package grammar

import (
	"fmt"
	"io"
)

func TypeSwiching(i interface{}) {
	switch j := i.(type) {
	case nil:
		fmt.Println("i는 nil, j타입은 interface{}: ", j)
	case int:
		fmt.Println("i 는 nil. j타입은 interface{}")
	case MyInt:
		fmt.Println("j타입은 MyInt")
	case io.Reader:
		fmt.Println("j타입은 io.Reader", j)
	case string:
		fmt.Println("j타입은 ", j)
	case bool, rune:
		fmt.Println("i타입은 bool 또는 rune 따라서 j타입은 interface{}", j)
	default:
		fmt.Println("i의 타입을 알 수 없으므로 j타입은 interface{}")
		fmt.Println("type:", j)
	}

}
