package conditionals

import "fmt"

func Conditionals() {
    // --- 1. 기본적인 if 문 ---

    // age 변수를 선언하고 초기값을 할당합니다.
    age := 20

    // age가 18보다 크거나 같을 경우에만 실행됩니다.
    // Go의 if 문은 조건식을 소괄호()로 감싸지 않습니다.
    if age >= 18 {
        fmt.Println("성인입니다.") // 조건이 참(true)일 경우 이 블록이 실행됩니다.
    }

    // --- 2. if-else 문 ---

    // isStudent 변수를 선언하고 초기값을 할당합니다.
    isStudent := true

    // isStudent가 참(true)일 경우와 거짓(false)일 경우를 나눕니다.
    if isStudent {
        fmt.Println("학생입니다.")
    } else {
        fmt.Println("학생이 아닙니다.") // isStudent가 거짓(false)일 경우 이 블록이 실행됩니다.
    }

    // --- 3. if-else if-else 문 (다중 조건) ---

    // score 변수를 선언하고 초기값을 할당합니다.
    score := 85

    // 여러 조건을 순서대로 확인합니다.
    // 위에서부터 아래로 조건을 확인하고, 참인 첫 번째 조건의 블록만 실행됩니다.
    if score >= 90 {
        fmt.Println("A 학점입니다.")
    } else if score >= 80 {
        // score가 90 미만이면서 80 이상일 경우 이 블록이 실행됩니다.
        fmt.Println("B 학점입니다.")
    } else if score >= 70 {
        // score가 80 미만이면서 70 이상일 경우 이 블록이 실행됩니다.
        fmt.Println("C 학점입니다.")
    } else {
        // 위의 모든 조건이 거짓(false)일 경우 이 블록이 실행됩니다.
        fmt.Println("F 학점입니다.")
    }
    
    // --- 4. if 문에 변수 선언 포함하기 ---

    // Go의 if 문은 조건식 앞에 짧은 선언(short declaration)을 포함할 수 있습니다.
    // 이 변수(result)는 if 문 블록과 else 블록 내부에서만 유효합니다.
    if result := 10 % 2; result == 0 {
        fmt.Println("10은 짝수입니다.")
    } else {
        fmt.Println("10은 홀수입니다.")
    }
    
    // 아래 코드는 컴파일 에러를 발생시킵니다.
    // fmt.Println(result) // 'result' 변수는 if-else 블록 외부에서 사용할 수 없습니다.

    // --- 5. 논리 연산자 사용하기 ---

    // && (AND): 두 조건이 모두 참일 때
    // || (OR): 두 조건 중 하나라도 참일 때
    // ! (NOT): 조건을 반대로 바꿀 때

    // grade 변수를 선언합니다.
    grade := "B"
    
    // grade가 "A"이거나 "B"일 경우를 확인합니다.
    if grade == "A" || grade == "B" {
        fmt.Println("잘했습니다!")
    } else {
        fmt.Println("노력합시다!")
    }

    // score가 80점 이상 90점 미만인지 확인합니다.
    if score >= 80 && score < 90 {
        fmt.Println("B학점 범위입니다.")
    }
}