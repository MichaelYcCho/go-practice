package grammar

import "fmt"

// 인터페이스는 Go의 유일한 추상 타입이다.

// 인터페이스는 이름의 마지막에 er를 붙인다.
// 인터페이스에의해 정의된 메서드는 인터페이스의 메서드 세트를 호출한다.

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// business logic
	fmt.Println("Data: ", data)
	return data
}

type Logic interface {
	Process(data string) string // 1. 인터페이스에서는 가지고 있는 메서드 세트는 반드시 존재해야한다.
}

type Client struct {
	L Logic // 2. Client는 Logic 인터페이스를 가지며 이는 반드시 Process 메서드를 가지고 있어야 한다.
}

func (c Client) Program() {
	// get data from somewhere
	c.L.Process("hi")

}

func InterfaceTest01() {
	c := Client{
		L: LogicProvider{}, // 3. LogicProvider의 Process가 없다면 에러가 발생한다.
		// 4. 이와는 별개로 LogicProvider는 선언된 것이 없는 빈 struct이다. 즉 모든 타입을 받을 수 있다.
	}
	c.Program()
}
