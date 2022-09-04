package goroutine

import "fmt"

func doSomething(x int) int {
	return x * 5
}

func RunSimpleGoroutine() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5) // 반환값을 취해 채널(ch)에 붙인다음
	}()
	fmt.Println(<-ch) // 그 채널에서 값을 끌어온다

}
