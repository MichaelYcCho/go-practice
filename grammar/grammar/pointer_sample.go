package grammar

import "fmt"

func Pointers() {
    // --- 1. 포인터 기초 개념 ---
    
    // 포인터는 메모리 주소를 저장하는 변수입니다.
    // 변수의 값 자체가 아니라, 그 값이 저장된 메모리 위치를 가리킵니다.
    // & 연산자: 변수의 주소를 가져옴 (address-of operator)
    // * 연산자: 포인터가 가리키는 값에 접근 (dereference operator)
    
    fmt.Println("--- 1. 포인터 기초 ---")
    
    // 일반 변수 선언
    age := 25
    fmt.Printf("age 변수의 값: %d\n", age)
    fmt.Printf("age 변수의 메모리 주소: %p\n", &age) // &age는 age 변수의 주소
    
    // 포인터 변수 선언
    var agePtr *int        // int형을 가리키는 포인터 선언 (초기값은 nil)
    agePtr = &age          // age 변수의 주소를 agePtr에 할당
    
    fmt.Printf("agePtr이 저장하는 주소: %p\n", agePtr)    // 포인터가 가리키는 주소 0x14000102020
    fmt.Printf("agePtr이 가리키는 값: %d\n", *agePtr)     // 포인터가 가리키는 실제 값 25
    
    // 짧은 선언으로 포인터 생성
    namePtr := &age        // age의 주소를 바로 할당
    fmt.Printf("namePtr이 가리키는 값: %d\n", *namePtr)
    
    // --- 2. 포인터를 통한 값 수정 ---
    
    fmt.Println("\n--- 2. 포인터를 통한 값 수정 ---")
    
    score := 100
    scorePtr := &score
    
    fmt.Printf("수정 전 score: %d\n", score) // 100
    fmt.Printf("수정 전 *scorePtr: %d\n", *scorePtr) // 100
    
    // 포인터를 통해 원본 값 수정
    *scorePtr = 95         // scorePtr이 가리키는 위치의 값을 95로 변경
    
    fmt.Printf("수정 후 score: %d\n", score)           // 원본 변수도 변경됨! 95
    fmt.Printf("수정 후 *scorePtr: %d\n", *scorePtr) // 95
    
    // --- 3. nil 포인터 ---
    
    fmt.Println("\n--- 3. nil 포인터 ---")
    
    var ptr *int           // 포인터 변수 선언 (초기값은 nil)
    fmt.Printf("ptr의 값: %v\n", ptr)                  // <nil>
    fmt.Printf("ptr이 nil인가? %t\n", ptr == nil)      // true
    
    // nil 포인터를 역참조하려고 하면 런타임 패닉 발생!
    // fmt.Println(*ptr)   // ❌ 이 줄의 주석을 해제하면 패닉 발생
    
    // 안전한 포인터 사용
    if ptr != nil {
        fmt.Printf("ptr이 가리키는 값: %d\n", *ptr)
    } else {
        fmt.Println("ptr은 nil이므로 역참조할 수 없습니다.")
    }
    
}