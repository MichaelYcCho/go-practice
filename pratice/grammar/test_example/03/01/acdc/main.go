// Package acdc asks if you are to rock
package acdc

// Sum adds an unlimited number of values of type int.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// go install golang.org/x/tools/cmd/godoc@latest
