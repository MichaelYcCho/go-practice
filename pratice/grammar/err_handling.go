package grammar

import (
	"errors"
	"fmt"
)

func ErrorHandling01(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("Odd number")
	}
	return i * 2, nil
}

func ErrorHandling02(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("Odd number: %d", i)
	}
	return i * 2, nil
}
