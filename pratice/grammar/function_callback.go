package grammar

import "fmt"

func sum01(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total

}

func CalculateNumber() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum01(ii...)
	fmt.Println(s)

	s2 := evenSum(sum01, ii...)
	fmt.Println(s2)
}

func evenSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)

}
