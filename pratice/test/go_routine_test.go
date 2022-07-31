package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoRoutine(t *testing.T) {

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
