package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoRoutine_01(t *testing.T) {

	t.Run("goroutine을 활용한 slice에 값 추가", func(t *testing.T) {
		var numbers [100]int
		var wg sync.WaitGroup
		wg.Add(100)
		for i := 0; i < 100; i++ {
			// Anonymous Function의 arf에 i를 전달해야 순차적으로 숫자가 input된다
			go func(i int) {
				// numbers slice에 i 추가
				//numbers = append(numbers, i) <- 완전히 랜덤값이 중복으로 들어오므로 fail
				numbers[i] = i

				//Done이 될때마다 Add 카운트는 -1이 된다
				wg.Done()
			}(i)
		}
		// wg의 카운트가 0이 될때까지 기다린다
		wg.Wait()

		// actual : [0 1 2 ... 99]
		var expected []int = make([]int, 100)
		// TODO expected를 만들어주세요.
		for i := 0; i < 100; i++ {
			expected[i] = i
		}

		assert.ElementsMatch(t, expected, numbers)
	})
}

func TestMutex(t *testing.T) {
	// Mutex(뮤텍스)는 여러 고루틴이 공유하는 데이터를 보호할 때 사용
	var mutex sync.Mutex
	var wg sync.WaitGroup

	var counter = 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			counter++

			defer wg.Done()
			defer mutex.Unlock()
		}()
	}
	wg.Wait()
	t.Logf("counter: %d", counter)
}
