package goroutine

import (
	"fmt"
	"math"
)

// 메소드 집합은 타입이 구현하는 인터페이스를 결정한다.

type methodCirlce struct {
	redius float64
}

type methodShape interface {
	area() float64
}

func (c *methodCirlce) area() float64 {
	return math.Pi * c.redius * c.redius
}

func info(s methodShape) {
	fmt.Println(s.area())
}

func RunMethodSet() {
	c := methodCirlce{10}
	//info(c)
	fmt.Println(c.area())
}
