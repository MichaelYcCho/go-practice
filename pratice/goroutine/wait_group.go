package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // sync패키지의 WaitGroup 타입

func RunTest01() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GOroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go test01()
	test02()

	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GOroutines\t", runtime.NumGoroutine())
	wg.Wait()

}

func test01() {
	for i := 0; i < 10; i++ {
		fmt.Println("test01", i)
	}
	wg.Done()
}

func test02() {
	for i := 0; i < 10; i++ {
		fmt.Println("test02", i)
	}

}
