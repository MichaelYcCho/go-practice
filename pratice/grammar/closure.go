package grammar

import (
	"fmt"
	"sort"
)

// 클로저는 함수내부에 선언된 함수를 의미한다.
// 내부에서 선언된 함수가 외부 함수에서 선언한 변수를 접근하고 수정할 수 있는 것을 의미한다.

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func Closure() []Person {

	people := []Person{
		{"Aohn", "Aoe", 30},
		{"Cane", "Coe", 25},
		{"Bary", "Boe", 20},
	}

	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println(people)
	return people
}
