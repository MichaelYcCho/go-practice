package grammar

import "fmt"

// 가변변수 핸들링

func ArgErrorCase() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}

func ArgSolveCase() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func(val int) {
			ch <- val * 2
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
