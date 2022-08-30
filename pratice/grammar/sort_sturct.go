package grammar

import (
	"fmt"
	"sort"
)

type SortPerson struct {
	First string
	Age   int
}

func (p SortPerson) String() string {
	return fmt.Sprintf("%s: %d", p.First, p.Age)
}

// type byAge[]sortPerson 으로 시작하는 코드 전체를 복사하여 핸들리이 가능하다
type byAge []SortPerson

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byName []SortPerson

func (bn byName) Len() int           { return len(bn) }
func (bn byName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func SortStruct() {
	p1 := SortPerson{"James", 20}
	p2 := SortPerson{"Miss", 27}
	p3 := SortPerson{"Q", 64}
	p4 := SortPerson{"M", 32}
	p5 := SortPerson{"A", 65}

	people := []SortPerson{p1, p2, p3, p4, p5}

	fmt.Println(people)

	fmt.Println("===============")
	sort.Sort(byAge(people))
	fmt.Println(people)

	fmt.Println("===============")
	sort.Sort(byName(people))
	fmt.Println(people)
}
