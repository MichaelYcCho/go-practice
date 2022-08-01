package grammar

import "fmt"

// 가변 파라미터
func Variadic_parameters(base int, vals ...int) []int {
	out := make([]int, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	fmt.Println(out)
	return out

}
