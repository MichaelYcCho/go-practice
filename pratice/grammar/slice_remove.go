package grammar

import "fmt"

func SliceRemove01() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 77, 88, 99)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}

// 마지막 요소를 명시한 index에 넣는법
func SliceRemove02() {
	slice := []string{"a", "b", "c", "d", "e", "f"}
	index := 1
	slice[index] = slice[len(slice)-1] // copy last element toindex we wish to remove
	slice[len(slice)-1] = ""           // remove the element
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
