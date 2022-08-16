package grammar

import (
	"fmt"
	"os"
)

// Go 런타임이 다음에 무슨일이 일어날지 알수 없을때 panic를 발생시킨다.
// panic후에는 defer()이 실행되며 defer가 main함수에 도달하면 프로그램은 종료된다.

func doPanic(msg string) {
	panic(msg)
}

func PanicTest() {
	doPanic(os.Args[0])
}

// recover 테스트 - recover이 일어나면 실행은 정상적으로 진행된다.

func div60(i int) {
	// recover은 잠재적인 panic을 처리하기 위해 defer 함수에서 호출한다.
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("예외 발생", v)
		}
	}()
	fmt.Println(60 / i)
}

func PanicTest02() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
