package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num of goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			// Done()은 해당 context를 위한 작업위 취소될때 닫힌 채널을 반환한다.
			case <-ctx.Done():
				return
			// 위의 상황이 아니면 하기의 로직처럼 0.2초 슬립에 빠졌다가 n을 출력한다.
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num of goroutines 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num of goroutines 3:", runtime.NumGoroutine())

}
