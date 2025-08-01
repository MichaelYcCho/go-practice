package func_func

import "fmt"

func AdvancedFunctions() {
    // --- 1. 가변 인자 함수 (Variadic Functions) ---

    // 가변 인자 함수는 정해지지 않은 수의 인자를 받을 수 있습니다.
    // 인자 타입 앞에 ...을 붙여서 선언합니다.
    fmt.Println("--- 1. 가변 인자 함수 ---")
    sum := calculateSum(1, 2, 3, 4, 5) // 5개의 인자를 전달
    fmt.Printf("1, 2, 3, 4, 5의 합계: %d\n", sum)

    anotherSum := calculateSum(10, 20) // 2개의 인자를 전달
    fmt.Printf("10, 20의 합계: %d\n", anotherSum)

    // --- 2. 익명 함수 (Anonymous Functions) ---

    // 이름이 없는 함수로, 주로 다른 함수의 인자로 전달하거나,
    // 특정 스코프 내에서 일회성으로 사용될 때 유용합니다.
    fmt.Println("\n--- 2. 익명 함수 ---")
    
    // 익명 함수를 변수에 할당
    greet := func(name string) {
        fmt.Printf("안녕하세요, %s님!\n", name)
    }
    greet("Go 개발자")

    // 익명 함수를 선언과 동시에 실행
    func() {
        fmt.Println("즉시 실행되는 익명 함수입니다.")
    }() // 괄호()를 붙여 즉시 실행

    // --- 3. 클로저 (Closures) ---

    // 클로저는 자신이 선언될 당시의 환경(변수)을 기억하는 함수입니다.
    // '외부 함수'의 변수에 '내부 함수'가 접근할 수 있게 해줍니다.
    fmt.Println("\n--- 3. 클로저 ---")
    
    // 클로저를 반환하는 함수를 호출하여 'increment'라는 함수를 받습니다.
    increment := createCounter()
    fmt.Printf("카운터: %d\n", increment()) // 1
    fmt.Printf("카운터: %d\n", increment()) // 2
    
    // 또 다른 클로저를 생성하여 완전히 독립적으로 동작합니다.
    increment2 := createCounter()
    fmt.Printf("새로운 카운터: %d\n", increment2()) // 1

    // --- 4. defer 키워드 ---

    // 'defer'는 해당 함수가 종료되기 직전에 실행될 함수를 지정합니다.
    // 주로 파일 닫기, 리소스 해제와 같은 뒷정리 작업에 사용됩니다.
    fmt.Println("\n--- 4. defer 키워드 ---")
    
    // 이 함수가 끝나기 직전에 fmt.Println("마지막에 실행됩니다.")가 실행됩니다.
    // defer 문은 LIFO(Last-In, First-Out) 방식으로 실행됩니다.
    defer fmt.Println("마지막에 실행됩니다.")
    
    fmt.Println("먼저 실행됩니다.")
}

// ---------------------- 함수 정의 ----------------------

// calculateSum 함수: 정해지지 않은 수의 int 인자를 받습니다.
// ...numbers는 'numbers'라는 이름의 int 슬라이스([]int)로 변환됩니다.
func calculateSum(numbers ...int) int {
    total := 0
    // range를 사용하여 슬라이스를 순회하며 합계를 계산합니다.
    for _, num := range numbers {
        total += num
    }
    return total
}

// createCounter 함수: 클로저를 반환합니다.
func createCounter() func() int {
    // 외부 함수 스코프에 있는 변수입니다.
    count := 0
    // 이 익명 함수가 'count' 변수를 기억합니다.
    return func() int {
        count++
        return count
    }
}