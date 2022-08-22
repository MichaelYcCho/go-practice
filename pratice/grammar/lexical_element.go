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

// 쏟아낸다는 뜻으로 공식적인 용어는 아님
func Unfurling() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...) // 이렇게 뒤에 ... 을 붙여서 int형들을 보내면, 아래의 sum에선 ...int를 통해 int를 다시 슬라이스로 받아온다
	fmt.Println(x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
