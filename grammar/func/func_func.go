package func_func

import "fmt"

func Func() {
    // --- 1. 매개변수와 반환값이 없는 함수 호출 ---

    // hello() 함수를 호출하여 코드를 실행합니다.
    fmt.Println("--- 1. 매개변수와 반환값이 없는 함수 호출 ---")
    hello()
    
    // --- 2. 매개변수만 있는 함수 호출 ---

    // sayHello 함수를 호출하며 "Go"라는 문자열을 전달합니다.
    fmt.Println("\n--- 2. 매개변수만 있는 함수 호출 ---")
    sayHello("Go")
    sayHello("Functions")

    // --- 3. 매개변수와 반환값이 모두 있는 함수 호출 ---

    // add 함수를 호출하고 반환된 값을 sum 변수에 저장합니다.
    // Go에서는 여러 개의 반환값을 가질 수 있습니다.
    fmt.Println("\n--- 3. 매개변수와 반환값이 모두 있는 함수 호출 ---")
    sum := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", sum)
    
    // multiply 함수는 두 개의 반환값을 가집니다.
    // area 변수에 첫 번째 반환값을, perimeter 변수에 두 번째 반환값을 저장합니다.
    area, perimeter := multiply(10, 5)
    fmt.Printf("가로 10, 세로 5 직사각형의 면적: %d, 둘레: %d\n", area, perimeter)
    
    // --- 4. 반환값 무시하기 ---
    
    // 만약 여러 반환값 중 특정 값을 사용하지 않을 경우, 밑줄(_)을 사용하여 무시할 수 있습니다.
    area2, _ := multiply(8, 4)
    fmt.Printf("\n--- 4. 반환값 중 하나만 사용 ---")
    fmt.Printf("\n가로 8, 세로 4 직사각형의 면적: %d\n", area2)
}

// ---------------------- 함수 정의 ----------------------

// hello 함수: 매개변수도, 반환값도 없습니다.
func hello() {
    fmt.Println("안녕하세요! 함수를 배우고 있어요.")
}

// sayHello 함수: 하나의 매개변수(name)를 받지만, 반환값은 없습니다.
func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// add 함수: 두 개의 매개변수(a, b)를 받고, 하나의 정수(int) 값을 반환합니다.
// 함수 시그니처: func 함수이름(매개변수명 타입, 매개변수명 타입) 반환타입
func add(a int, b int) int {
    return a + b
}

// multiply 함수: 두 개의 매개변수를 받고, 두 개의 정수(int) 값을 반환합니다.
// 반환값이 여러 개일 경우, 반환 타입을 소괄호로 묶습니다.
func multiply(width int, height int) (int, int) {
    area := width * height
    perimeter := 2 * (width + height)
    return area, perimeter
}