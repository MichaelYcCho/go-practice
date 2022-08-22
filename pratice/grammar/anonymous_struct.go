package grammar

import "fmt"

// 일반적인 구조체
type person = struct {
	first string
	last  string
	age   int
}

func StructExample() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   35,
	}
	fmt.Println(p1)

}

func AnonymousStruct() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   35,
	}
	fmt.Println(p1)
}
