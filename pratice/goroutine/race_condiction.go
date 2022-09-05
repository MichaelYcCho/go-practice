package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func RaceFunction() {
	fmt.Println("CPU :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()                                    // 보다 안전한 방식으로 counter에 write한다.
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) // 보다 안전한 방식으로 counter에 read한다.

			wg.Done()
		}()
		fmt.Println("Goroutine :", runtime.NumGoroutine()) // Gosched는 다른 goroutine이 실행하게 하면스 그 프로세서를 양보한다.
	}
	wg.Wait()

	fmt.Println("Goroutine :", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
