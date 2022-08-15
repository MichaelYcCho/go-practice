package grammar

import (
	"errors"
	"fmt"
)

func ErrorHandling01(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("odd number")
	}
	return i * 2, nil
}

func ErrorHandling02(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("odd number: %d", i)
	}
	return i * 2, nil
}
