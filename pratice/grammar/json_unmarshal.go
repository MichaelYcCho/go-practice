package grammar

import (
	"encoding/json"
	"fmt"
)

type unMarhshalPerson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func JsonUnMarshal() {
	s := `[{"First":"James","Last":"Bond","Age":20},{"First":"Miss","Last":"Moneypenny","Age":19}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//var people := []unMarhshalPerson{}
	var people []unMarhshalPerson

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", people)

}
