package grammar

import "fmt"

func ArraysAndSlices() {
    // --- 1. 배열(Array) 기초 ---
    
    // 배열은 고정된 크기를 가지는 같은 타입 요소들의 연속된 메모리 공간입니다.
    // 선언: var 변수명 [크기]타입
    // 배열의 크기는 타입의 일부이므로, [3]int와 [5]int는 다른 타입입니다.
    
    fmt.Println("--- 1. 배열(Array) 선언과 초기화 ---")
    
    // 1-1. 배열 선언 (제로 값으로 초기화)
    var numbers [5]int // 크기가 5인 정수 배열, 모든 요소가 0으로 초기화
    fmt.Printf("초기화된 배열: %v\n", numbers) // [0 0 0 0 0]
    
    // 1-2. 배열 선언과 동시에 초기화
    fruits := [3]string{"apple", "banana", "orange"}
    fmt.Printf("과일 배열: %v\n", fruits) // [apple banana orange]
    
    // 1-3. 컴파일러가 길이를 자동으로 계산하게 하기
    colors := [...]string{"red", "green", "blue", "yellow"}
    fmt.Printf("색깔 배열: %v (길이: %d)\n", colors, len(colors)) // 길이가 4로 자동 설정
    
    // --- 2. 배열 요소 접근 및 수정 ---
    
    fmt.Println("\n--- 2. 배열 요소 접근 및 수정 ---")
    
    // 인덱스를 사용하여 배열 요소에 접근합니다 (0부터 시작)
    fmt.Printf("첫 번째 과일: %s\n", fruits[0]) // "apple"
    fmt.Printf("두 번째 과일: %s\n", fruits[1]) // "banana"
    
    // 배열 요소 값 변경
    fruits[1] = "mango"
    fmt.Printf("변경 후 과일 배열: %v\n", fruits) // [apple mango orange]
    
    // 배열 길이 확인
    fmt.Printf("배열의 길이: %d\n", len(fruits))
    
    // --- 3. 배열 순회하기 ---
    
    fmt.Println("\n--- 3. 배열 순회하기 ---")
    
    // 3-1. 일반 for 루프 사용
    fmt.Println("일반 for 루프:")
    for i := 0; i < len(colors); i++ {
        fmt.Printf("인덱스 %d: %s\n", i, colors[i])
    }
    
    // 3-2. range 키워드 사용 (더 관용적인 방법)
    fmt.Println("\nrange 키워드 사용:")
    for index, value := range colors {
        fmt.Printf("인덱스 %d: %s\n", index, value)
    }
    
    // --- 4. 슬라이스(Slice) 기초 ---
    
    // 슬라이스는 배열의 유연한 버전입니다. 크기가 동적으로 변할 수 있습니다.
    // 슬라이스는 실제로는 배열을 가리키는 포인터, 길이, 용량 정보를 담는 구조체입니다.
    
    fmt.Println("\n--- 4. 슬라이스(Slice) 선언과 초기화 ---")
    
    // 4-1. 빈 슬라이스 선언
    var emptySlice []int
    fmt.Printf("빈 슬라이스: %v (길이: %d, 용량: %d)\n", emptySlice, len(emptySlice), cap(emptySlice))
    
    // 4-2. make 함수를 사용하여 슬라이스 생성
    // make([]타입, 길이, 용량)
    slice1 := make([]int, 3)    // 길이 3, 용량 3인 슬라이스
    slice2 := make([]int, 3, 5) // 길이 3, 용량 5인 슬라이스
    fmt.Printf("slice1: %v (길이: %d, 용량: %d)\n", slice1, len(slice1), cap(slice1))
    fmt.Printf("slice2: %v (길이: %d, 용량: %d)\n", slice2, len(slice2), cap(slice2))
    
    // 4-3. 슬라이스 리터럴로 초기화
    animals := []string{"dog", "cat", "bird"}
    fmt.Printf("동물 슬라이스: %v (길이: %d, 용량: %d)\n", animals, len(animals), cap(animals))
    
    // --- 5. 슬라이스 조작 ---
    
    fmt.Println("\n--- 5. 슬라이스 조작 ---")
    
    // 5-1. append 함수로 요소 추가
    // append는 새로운 슬라이스를 반환하므로 결과를 다시 할당해야 합니다
    animals = append(animals, "fish")
    fmt.Printf("요소 추가 후: %v\n", animals)
    
    // 여러 요소를 한 번에 추가
    animals = append(animals, "rabbit", "hamster")
    fmt.Printf("여러 요소 추가 후: %v\n", animals)
    
    // 다른 슬라이스의 모든 요소를 추가 (... 연산자 사용)
    moreAnimals := []string{"tiger", "lion"}
    animals = append(animals, moreAnimals...)
    fmt.Printf("슬라이스 병합 후: %v\n", animals)
    
    // --- 6. 슬라이싱 (Slicing) ---
    
    // 기존 배열이나 슬라이스에서 일부분을 추출하여 새로운 슬라이스를 만드는 것
    // 문법: slice[시작인덱스:끝인덱스] (끝인덱스는 포함하지 않음)
    
    fmt.Println("\n--- 6. 슬라이싱 (Slicing) ---")
    
    numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Printf("원본 슬라이스: %v\n", numbers2)
    
    // 인덱스 2부터 5까지 (5는 포함하지 않음)
    subSlice1 := numbers2[2:5]
    fmt.Printf("numbers2[2:5]: %v\n", subSlice1) // [2 3 4]
    
    // 처음부터 인덱스 4까지
    subSlice2 := numbers2[:4]
    fmt.Printf("numbers2[:4]: %v\n", subSlice2) // [0 1 2 3]
    
    // 인덱스 6부터 끝까지
    subSlice3 := numbers2[6:]
    fmt.Printf("numbers2[6:]: %v\n", subSlice3) // [6 7 8 9]
    
    // 전체 복사
    subSlice4 := numbers2[:]
    fmt.Printf("numbers2[:]: %v\n", subSlice4) // [0 1 2 3 4 5 6 7 8 9]
    
    // --- 7. 슬라이스는 참조 타입 ---
    
    fmt.Println("\n--- 7. 슬라이스는 참조 타입 ---")
    
    // 슬라이싱으로 만든 슬라이스는 원본과 같은 배열을 참조합니다
    original := []int{1, 2, 3, 4, 5}
    reference := original[1:4] // [2 3 4]
    
    fmt.Printf("원본: %v\n", original)
    fmt.Printf("참조: %v\n", reference)
    
    // 참조 슬라이스의 값을 변경하면 원본도 영향을 받습니다
    reference[0] = 99
    fmt.Printf("참조 변경 후 원본: %v\n", original) // [1 99 3 4 5]
    fmt.Printf("참조 변경 후 참조: %v\n", reference) // [99 3 4]
    
    // --- 8. 슬라이스 복사 (copy 함수) ---
    
    fmt.Println("\n--- 8. 슬라이스 복사 ---")
    
    // copy 함수를 사용하여 완전히 독립적인 복사본을 만들 수 있습니다
    source := []int{10, 20, 30, 40, 50}
    destination := make([]int, len(source))
    
    // copy(목적지, 원본)
    copied := copy(destination, source)
    fmt.Printf("복사된 요소 개수: %d\n", copied)
    fmt.Printf("원본: %v\n", source)
    fmt.Printf("복사본: %v\n", destination)
    
    // 복사본 변경이 원본에 영향을 주지 않음
    destination[0] = 999
    fmt.Printf("복사본 변경 후 원본: %v\n", source)      // [10 20 30 40 50] (변경 없음)
    fmt.Printf("복사본 변경 후 복사본: %v\n", destination) // [999 20 30 40 50]
    
    // --- 9. 다차원 슬라이스 ---
    
    fmt.Println("\n--- 9. 다차원 슬라이스 ---")
    
    // 슬라이스의 슬라이스를 만들어 2차원 배열처럼 사용할 수 있습니다
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("2차원 슬라이스:")
    for i, row := range matrix {
        fmt.Printf("행 %d: %v\n", i, row)
    }
    
    // 개별 요소 접근
    fmt.Printf("matrix[1][2] = %d\n", matrix[1][2]) // 6
    
    // --- 10. 실용적인 예제: 슬라이스로 스택 구현 ---
    
    fmt.Println("\n--- 10. 슬라이스로 스택 구현 ---")
    
    var stack []int
    
    // Push 연산 (요소 추가)
    stack = append(stack, 10)
    stack = append(stack, 20)
    stack = append(stack, 30)
    fmt.Printf("스택에 요소 추가 후: %v\n", stack)
    
    // Pop 연산 (마지막 요소 제거 및 반환)
    if len(stack) > 0 {
        top := stack[len(stack)-1]        // 마지막 요소 가져오기
        stack = stack[:len(stack)-1]      // 마지막 요소 제거
        fmt.Printf("Pop된 요소: %d\n", top)
        fmt.Printf("Pop 후 스택: %v\n", stack)
    }
}

// --- 추가 유틸리티 함수들 ---

// 슬라이스에서 특정 값 찾기
func findInSlice(slice []string, target string) int {
    for i, value := range slice {
        if value == target {
            return i // 찾은 인덱스 반환
        }
    }
    return -1 // 찾지 못했으면 -1 반환
}

// 슬라이스에서 특정 인덱스의 요소 제거
func removeAtIndex(slice []int, index int) []int {
    if index < 0 || index >= len(slice) {
        return slice // 잘못된 인덱스면 원본 반환
    }
    
    // 인덱스 앞부분과 뒷부분을 합쳐서 새로운 슬라이스 생성
    return append(slice[:index], slice[index+1:]...)
}

func SliceUtilities() {
    fmt.Println("\n=== 슬라이스 유틸리티 함수 예제 ===")
    
    // findInSlice 함수 테스트
    fruits := []string{"apple", "banana", "orange", "grape"}
    index := findInSlice(fruits, "orange")
    if index != -1 {
        fmt.Printf("'orange'을 인덱스 %d에서 찾았습니다.\n", index)
    } else {
        fmt.Println("'orange'을 찾지 못했습니다.")
    }
    
    // removeAtIndex 함수 테스트
    numbers := []int{10, 20, 30, 40, 50}
    fmt.Printf("원본: %v\n", numbers)
    
    numbers = removeAtIndex(numbers, 2) // 인덱스 2 (값 30) 제거
    fmt.Printf("인덱스 2 제거 후: %v\n", numbers)
}