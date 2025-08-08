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
    
    // --- 4. new 함수로 포인터 생성 ---
    
    fmt.Println("\n--- 4. new 함수로 포인터 생성 ---")
    
    // new 함수는 지정된 타입의 제로 값으로 초기화된 메모리를 할당하고
    // 그 메모리의 주소를 반환합니다.
    numPtr := new(int)     // int 타입의 제로 값(0)으로 초기화된 메모리 할당
    fmt.Printf("new로 생성한 포인터: %p\n", numPtr)
    fmt.Printf("new로 생성한 포인터가 가리키는 값: %d\n", *numPtr) // 0
    
    *numPtr = 42           // 값 할당
    fmt.Printf("값 할당 후: %d\n", *numPtr)
    
    // --- 5. 함수와 포인터 ---
    
    fmt.Println("\n--- 5. 함수와 포인터 ---")
    
    x := 10
    y := 20
    
    fmt.Printf("함수 호출 전 - x: %d, y: %d\n", x, y)
    
    // 값 전달 (Call by Value)
    swapByValue(x, y)
    fmt.Printf("swapByValue 호출 후 - x: %d, y: %d\n", x, y) // 값이 바뀌지 않음!
    
    // 포인터 전달 (Call by Reference)
    swapByPointer(&x, &y)
    fmt.Printf("swapByPointer 호출 후 - x: %d, y: %d\n", x, y) // 값이 바뀜!
    
    // --- 6. 구조체와 포인터 ---
    
    fmt.Println("\n--- 6. 구조체와 포인터 ---")
    
    // PersonInfo 구조체 정의는 아래에 있습니다
    person1 := PersonInfo{Name: "Alice", Age: 30}
    fmt.Printf("person1: %+v\n", person1)
    
    // 구조체의 포인터 생성
    personPtr := &person1
    fmt.Printf("personPtr이 가리키는 구조체: %+v\n", *personPtr)
    
    // 포인터를 통한 필드 접근 (자동 역참조)
    // Go는 (*personPtr).Name 대신 personPtr.Name으로 편리하게 접근 가능
    fmt.Printf("이름: %s\n", personPtr.Name)     // (*personPtr).Name과 동일
    fmt.Printf("나이: %d\n", personPtr.Age)      // (*personPtr).Age와 동일
    
    // 포인터를 통한 필드 수정
    personPtr.Age = 31
    fmt.Printf("나이 수정 후 person1: %+v\n", person1) // 원본이 수정됨
    
    // --- 7. 슬라이스와 포인터 ---
    
    fmt.Println("\n--- 7. 슬라이스와 포인터 ---")
    
    // 슬라이스는 이미 참조 타입이지만, 슬라이스 자체의 포인터도 만들 수 있습니다
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("원본 슬라이스: %v\n", numbers)
    
    // 슬라이스의 포인터
    numbersPtr := &numbers
    fmt.Printf("포인터를 통한 슬라이스 접근: %v\n", *numbersPtr)
    
    // 슬라이스 요소의 포인터
    firstElementPtr := &numbers[0]
    fmt.Printf("첫 번째 요소의 주소: %p\n", firstElementPtr)
    fmt.Printf("첫 번째 요소의 값: %d\n", *firstElementPtr)
    
    *firstElementPtr = 100
    fmt.Printf("첫 번째 요소 수정 후: %v\n", numbers)
    
    // --- 8. 포인터 배열과 배열의 포인터 ---
    
    fmt.Println("\n--- 8. 포인터 배열과 배열의 포인터 ---")
    
    a, b, c := 1, 2, 3
    
    // 포인터들의 배열 (포인터 배열)
    ptrArray := []*int{&a, &b, &c}
    fmt.Println("포인터 배열:")
    for i, ptr := range ptrArray {
        fmt.Printf("  ptrArray[%d]이 가리키는 값: %d\n", i, *ptr)
    }
    
    // 배열의 포인터
    intArray := [3]int{10, 20, 30}
    arrayPtr := &intArray
    fmt.Printf("배열의 포인터가 가리키는 배열: %v\n", *arrayPtr)
    
    // 배열 포인터를 통한 요소 접근
    (*arrayPtr)[0] = 100    // 명시적 역참조
    arrayPtr[1] = 200       // Go의 자동 역참조 (편의 기능)
    fmt.Printf("수정 후 배열: %v\n", intArray)
    
    // --- 9. 포인터 체인 (포인터의 포인터) ---
    
    fmt.Println("\n--- 9. 포인터 체인 ---")
    
    value := 42
    ptr1 := &value          // value의 포인터
    ptr2 := &ptr1           // ptr1의 포인터 (이중 포인터)
    
    fmt.Printf("value: %d\n", value)
    fmt.Printf("ptr1이 가리키는 값: %d\n", *ptr1)
    fmt.Printf("ptr2가 가리키는 포인터가 가리키는 값: %d\n", **ptr2)
    
    // 이중 포인터를 통한 값 수정
    **ptr2 = 99
    fmt.Printf("이중 포인터로 수정 후 value: %d\n", value)
    
    // --- 10. 실용적인 예제: 연결 리스트 노드 ---
    
    fmt.Println("\n--- 10. 연결 리스트 노드 예제 ---")
    
    // 간단한 연결 리스트 노드 구조
    demonstrateLinkedList()
}

