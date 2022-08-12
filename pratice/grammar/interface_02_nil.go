package grammar

import "fmt"

// 인터페이스가 nil이되려면 타입과 값 모두 nil이여야한다.

func InterfaceNilTest() {
	var s *string
	fmt.Println(s == nil) // true

	var i interface{}
	fmt.Println(i == nil) // true

	i = s
	fmt.Println(i == nil) // false

}
