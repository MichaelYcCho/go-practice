package variable

import "fmt" // fmt 패키지는 입출력(format) 관련 함수들을 제공합니다.

func Variable() { 
    // --- 1. var 키워드를 사용한 변수 선언 ---

    // 1-1. 타입과 함께 선언 (초기값 없음):
    // Go는 변수를 선언할 때 타입을 명시적으로 지정할 수 있습니다.
    // 초기값을 지정하지 않으면 각 타입의 '제로 값(zero value)'으로 자동 초기화됩니다.
    // - int: 0
    // - float64: 0.0
    // - bool: false
    // - string: "" (빈 문자열)
    var age int // 'age'라는 이름의 정수(int)형 변수 선언. 초기값은 0
    var name string // 'name'이라는 이름의 문자열(string)형 변수 선언. 초기값은 ""
    var isGoLearner bool // 'isGoLearner'라는 이름의 불린(bool)형 변수 선언. 초기값은 false

    fmt.Println("--- 1-1. var 키워드를 사용한 변수 선언 (초기값 없음) ---")
    fmt.Printf("age: %d\n", age) // %d는 정수 값을 출력하는 서식 지정자입니다.
    fmt.Printf("name: \"%s\"\n", name) // %s는 문자열 값을 출력하는 서식 지정자입니다. \"는 큰따옴표를 문자열 안에 포함시킵니다.
    fmt.Printf("isGoLearner: %t\n", isGoLearner) // %t는 불린 값을 출력하는 서식 지정자입니다.
    fmt.Println("--------------------------------------------------")

    // 1-2. 타입과 함께 선언 (초기값 지정):
    // 선언과 동시에 초기값을 할당할 수 있습니다.
    var height float64 = 175.5 // 'height'라는 이름의 실수(float64)형 변수를 175.5로 초기화
    var city string = "Seoul" // 'city'라는 이름의 문자열(string)형 변수를 "Seoul"로 초기화

    fmt.Println("\n--- 1-2. var 키워드를 사용한 변수 선언 (초기값 지정) ---")
    fmt.Printf("height: %.1f\n", height) // %.1f는 소수점 이하 첫째 자리까지 출력하는 실수 서식 지정자입니다.
    fmt.Printf("city: %s\n", city)
    fmt.Println("--------------------------------------------------")

    // 1-3. 초기값만 지정 (타입 추론):
    // 초기값을 지정하면 Go 컴파일러가 초기값의 타입을 보고 변수의 타입을 자동으로 추론합니다.
    var country = "South Korea" // 초기값이 문자열이므로, 'country'는 string 타입으로 추론됩니다.
    var year = 2025 // 초기값이 정수이므로, 'year'는 int 타입으로 추론됩니다.

    fmt.Println("\n--- 1-3. 초기값만 지정 (타입 추론) ---")
    fmt.Printf("country: %s (타입: %T)\n", country, country) // %T는 변수의 타입을 출력하는 서식 지정자입니다.
    fmt.Printf("year: %d (타입: %T)\n", year, year)
    fmt.Println("--------------------------------------------------")

    // --- 2. 짧은 선언 (Short Declaration Operator): := ---

    // 이 방식은 함수 내부에서만 사용할 수 있습니다.
    // 변수 선언과 동시에 초기값을 할당하며, 타입은 Go가 자동으로 추론합니다.
    // 기존에 선언된 변수에는 사용할 수 없고, 새로운 변수에만 사용할 수 있습니다.
    
    // 이 선언 방식을 가장 흔하게 사용하게 될 것입니다.
    message := "Go 언어는 정말 재미있어요!" // 'message'는 string 타입으로 추론
    version := 1.22 // 'version'은 float64 타입으로 추론

    fmt.Println("\n--- 2. 짧은 선언 (Short Declaration Operator: :=) ---")
    fmt.Printf("message: %s\n", message)
    fmt.Printf("version: %.2f\n", version)
    fmt.Println("--------------------------------------------------")

    // --- 3. 여러 변수 동시 선언 ---

    // 여러 변수를 한 줄에 선언하고 초기화할 수 있습니다.
    var x, y, z int = 10, 20, 30 // x, y, z 모두 int 타입으로 선언 및 초기화
    
    // 타입 추론을 사용하여 여러 변수 동시 선언 및 초기화
    a, b := "hello", "world" // a는 string, b도 string으로 추론

    fmt.Println("\n--- 3. 여러 변수 동시 선언 ---")
    fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)
    fmt.Printf("a: %s, b: %s\n", a, b)
    fmt.Println("--------------------------------------------------")

    // --- 4. 변수 값 변경 ---

    // Go에서는 변수 선언 후 값을 자유롭게 변경할 수 있습니다.
    // 단, 선언된 변수의 타입과 일치하는 값만 할당할 수 있습니다.
    score := 100 // score를 100으로 초기화
    fmt.Printf("\n초기 score: %d\n", score)

    score = 95 // score 값을 95로 변경 (같은 int 타입)
    fmt.Printf("변경 후 score: %d\n", score)

    // score = "bad" // 에러 발생! int 타입 변수에 string 값을 할당할 수 없습니다.

    fmt.Println("--------------------------------------------------")

    // --- 5. 상수 선언 (Constants) ---
    // const 키워드를 사용하여 상수를 선언할 수 있습니다.
    // 상수는 선언 후에는 값을 변경할 수 없습니다.
    const PI float64 = 3.14159 // PI는 float64 타입의 상수
    const Greeting = "안녕하세요!" // Greeting은 string 타입의 상수

    fmt.Println("\n--- 5. 상수 선언 (Constants) ---")
    fmt.Printf("PI 값: %.5f\n", PI)
    fmt.Printf("Greeting 메시지: %s\n", Greeting)
    // PI = 3.0 // 에러 발생! 상수의 값은 변경할 수 없습니다.
    fmt.Println("--------------------------------------------------")
}