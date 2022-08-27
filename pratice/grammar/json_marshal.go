package grammar

import (
	"encoding/json"
	"fmt"
)

type marshalPerson struct {
	First string
	Last  string
	Age   int
}

func JsonMarshal() {
	p1 := marshalPerson{
		First: "James",
		Last:  "Bond",
		Age:   20,
	}
	p2 := marshalPerson{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   19,
	}

	people := []marshalPerson{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
