package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func RaceFunction() {
	fmt.Println("CPU :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() // Gosched는 다른 goroutine이 실행하게 하면스 그 프로세서를 양보한다.
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Goroutine :", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
