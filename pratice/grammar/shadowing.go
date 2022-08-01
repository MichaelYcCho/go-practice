package grammar

import "fmt"

func Shadowing() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5 // x를 섀도잉 민액 x = 5얐디먄 최상단 x는 5로 출력된다.
		fmt.Println(x)
	}
	fmt.Println(x)
	return
}
