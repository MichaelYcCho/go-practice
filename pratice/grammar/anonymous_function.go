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
