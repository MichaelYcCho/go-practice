package grammar

import "fmt"

func Pointer_01() {
	x := 10
	pointerToX := &x

	fmt.Println(pointerToX)  // 메모리 주소출력
	fmt.Println(*pointerToX) // 메모리 주소에 저장된 값 출력  -> 10

	z := 5 + *pointerToX
	fmt.Println(z) // 15
}
