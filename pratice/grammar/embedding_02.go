package grammar

import "fmt"

type Inner02 struct {
	A int
}

func (i Inner02) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner02) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer02 struct {
	Inner02
	S string
}

func (o Outer02) IntPrinter(val int) string {
	return fmt.Sprintf("Outer02: %d", val)
}

func (o Outer02) Double() string {
	return o.IntPrinter(o.Inner02.A * 2)
}

func EmbeddingStart03() {
	o := Outer02{
		Inner02: Inner02{
			A: 10,
		},
		S: "Hello",
	}

	/*
		임베딩된 항목에서 다른 메서드를 호출하는 임베딩된 항목의 메서드가 있고 (func (i Inner02) IntPrinter),
		포함하는 구조체에서 같은이름의 메서드가 있다면 (func (o Outer02) IntPrinter)
		이 메서드는 임베딩된 항목의 메서드를 호출한다.
		그러므로 임베딩을 무조건적인 상속의 개념으로 보면 안된다.

		물론 func (i Inner02) IntPrinter가 없다면 func (o Outer02) IntPrinter를 호출한다.

	*/
	fmt.Println(o.Inner02.A)      // 10
	fmt.Println(o.Double())       // Inner: 20
	fmt.Println(o.IntPrinter(10)) // Outer02: 10
}
