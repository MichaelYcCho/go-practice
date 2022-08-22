package grammar

import "fmt"

func ManyInt() {
	int_list(1, 2, 3, 4, 5)
}

// 가변매개인자(Lexical element)
/// 타입부분앞에 ...을 사용하는것으로 여러개의 매개인자를 받을 수 있다.
func int_list(args ...int) {
	fmt.Println(args)
	fmt.Printf("%T\n", args)
}
