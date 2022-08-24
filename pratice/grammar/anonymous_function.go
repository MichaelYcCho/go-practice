package grammar

import "fmt"

// 익명함수는 func 바로뒤에 입력파라미터와 반환값을 넣고 {}를 사용하여 선언한다.
// 아래의 메서드는 i변수를 내부익명함수로 전달하여 입력파라미터인 j로 할당하는 구조이다.
func AnonymousFunc() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("i값: ", i)
			fmt.Println("j값: ", j)
			fmt.Println()

		}(i)
	}
}

// 일급객체 표현식
func FunctionExpression() {
	f := func(x int) {
		fmt.Println("First func expression", x)
	}
	f(42)
}

func ReturnString() string {
	return "Hello World"
}

// 함수를 반환하는것도 가능하다
func ReturnFunction() {
	s1 := ReturnString()
	fmt.Println(s1)

}

// 함수에서 함수를 반환하는것도 가능하다
func ReturnFunction2() func() int {
	return func() int {
		return 42
	}

}
