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

type Foo struct {
	Field1 string
	Field2 int
}

// 함수로 구조체 전달을 포인터로 하여 항목을 채우는것보다

func Wrongway(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// 함수 내에서 구조체를 초기화 하고 반환하는 것이 좋다
func Rightway(Foo) (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}
