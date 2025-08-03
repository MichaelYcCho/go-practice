package main

import (
	"fmt"

	"github.com/yc/go-practice/practice/grammar"
)

// grammar 패키지를 가져옴

func main() {
    // CreateCounter 함수를 호출하여 첫 번째 카운터 클로저를 만듭니다.
    // 'counter1'은 func() int 타입의 함수가 됩니다.
    // 이 함수는 CreateCounter 내부의 'count' 변수를 기억합니다.
    counter1 := grammar.CreateCounter()

    fmt.Println("--- 첫 번째 카운터 ---")
    fmt.Println(counter1()) // 1
    fmt.Println(counter1()) // 2
    fmt.Println(counter1()) // 3

    // CreateCounter 함수를 다시 호출하여 두 번째 카운터 클로저를 만듭니다.
    // 'counter2'는 'counter1'과는 완전히 독립적인 'count' 변수를 기억합니다.
    counter2 := grammar.CreateCounter()

    fmt.Println("\n--- 두 번째 카운터 (독립적) ---")
    fmt.Println(counter2()) // 1
    fmt.Println(counter2()) // 2
    
    // 다시 첫 번째 카운터를 호출하면, 기존의 count 변수 값을 계속 이어서 사용합니다.
    fmt.Println("\n--- 첫 번째 카운터 다시 사용 ---")
    fmt.Println(counter1()) // 4 (이전에 3까지 세었기 때문에 4가 나옴)
}
