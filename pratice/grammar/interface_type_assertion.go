package grammar

import "fmt"

// type assertion - 타입단언
// 인터페이스를 구현한 구체 타입의 이름을 지정하거나
// 인터페이스 기반인 구체 타입에 의해 구현된 다른 인터페이스의 이름을 지정

type MyInt int

func TypeAssertionTest01() {
	var i interface{}
	var mine MyInt = 20
	i = mine // i 는 MyInt 타입
	i2 := i.(MyInt)
	// i2 := i.(string)
	// 타입단언은 기본 값의 타입과 반드시 일치해야함.
	// panic: interface conversion: interface {} is grammar.MyInt, not string
	fmt.Println(i2) // 20

}

func TypeAssertionTest02() error {
	var i interface{}
	var mine MyInt = 20
	i = mine // i 는 MyInt 타입
	i2, ok := i.(int)
	if !ok {
		// 여기 분기를 탐
		return fmt.Errorf("unexpected type for %v", i)
	}
	fmt.Println(i2 + 1) // 20
	return nil
}
