package grammar

import (
	"fmt"
	"sort"
)

func SortSlice01() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("======================")
	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("======================")
	sort.Sort(sort.Reverse(sort.IntSlice(xi)))
	sort.Sort(sort.Reverse(sort.StringSlice(xs)))

	fmt.Println(xi)
	fmt.Println(xs)

}
