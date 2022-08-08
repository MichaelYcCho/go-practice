package grammar

import (
	"fmt"
	"time"
)

/*
- 메서드가 리시버를 수정해야 한다면 반드시 포인터리시버를 사용해야한다.

*/

type Counter struct {
	total       int
	lastUpdated time.Time
}

// 리시버에 대한 수정이 필요한 경우 포인터를 사용해야한다.
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("%d", c.total)
}

func MethodTest() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}
