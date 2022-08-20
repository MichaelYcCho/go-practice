package grammar

import "fmt"

// 교착상태의 고틴
/*
Go 런타임에서 처음 시작할 때 동작하는 고루틴에서 실행된다.
죽 ch1을 읽을 때까지 계속 진행할 수 없고 main 고루틴은 ch2를 읽을 떄까지 진행할 수 없게된다.
*/
func DeadLockCase() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	ch2 <- v
	v2 := <-ch1
	fmt.Println(v, v2)

}

func DeadLockSolveCase() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:

	}

	fmt.Println(v, v2)
}
