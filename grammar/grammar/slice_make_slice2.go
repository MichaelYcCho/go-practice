package grammar

import "fmt"

/*
용량을 초과하면 Go가 자동으로 다음과 같은 일을 합니다:

자동 확장: 더 큰 새로운 배열을 메모리에 할당
데이터 복사: 기존 모든 요소를 새 배열로 복사
요소 추가: 새 요소를 추가
참조 업데이트: 새 배열을 가리키는 슬라이스 반환

용량 증가 규칙
Go의 슬라이스 용량 증가는 다음 규칙을 따릅니다:

작은 슬라이스 (< 1024): 약 2배씩 증가
큰 슬라이스 (≥ 1024): 약 1.25배씩 증가

*/

func SliceCapacityTwo() {
    fmt.Println("=== 슬라이스 용량 초과 테스트 ===")
    
    // 길이 3, 용량 5인 슬라이스 생성
    slice := make([]int, 3, 5)
    fmt.Printf("초기 상태: %v (길이: %d, 용량: %d)\n", slice, len(slice), cap(slice))
    
    // 초기값 설정
    slice[0] = 10
    slice[1] = 20
    slice[2] = 30
    fmt.Printf("값 설정 후: %v (길이: %d, 용량: %d)\n", slice, len(slice), cap(slice))
    
    fmt.Println("\n--- append로 요소 추가하기 ---")
    
    // 첫 번째 append: 용량 내에서 추가
    slice = append(slice, 40)
    fmt.Printf("40 추가 후: %v (길이: %d, 용량: %d)\n", slice, len(slice), cap(slice))
    
    // 두 번째 append: 용량 내에서 추가
    slice = append(slice, 50)
    fmt.Printf("50 추가 후: %v (길이: %d, 용량: %d)\n", slice, len(slice), cap(slice))
    
    // 세 번째 append: 용량 초과! 
    fmt.Println("\n🚨 용량 초과 상황!")
    slice = append(slice, 60)
    fmt.Printf("60 추가 후: %v (길이: %d, 용량: %d)\n", slice, len(slice), cap(slice))
    
    // 네 번째 append: 새로운 용량에서 추가
    slice = append(slice, 70)
    fmt.Printf("70 추가 후: %v (길이: %d, 용량: %d)\n", slice, len(slice), cap(slice))
    
    fmt.Println("\n=== 용량 증가 패턴 관찰 ===")
    
    // 작은 슬라이스부터 시작해서 용량 변화 관찰
    smallSlice := make([]int, 0, 1) // 길이 0, 용량 1
    fmt.Printf("시작: %v (길이: %d, 용량: %d)\n", smallSlice, len(smallSlice), cap(smallSlice))
    
    for i := 1; i <= 20; i++ {
        oldCap := cap(smallSlice)
        smallSlice = append(smallSlice, i)
        newCap := cap(smallSlice)
        
        if newCap != oldCap {
            fmt.Printf("요소 %d 추가 시 용량 변경: %d → %d (약 %0.1f배)\n", 
                i, oldCap, newCap, float64(newCap)/float64(oldCap))
        }
    }
    
    fmt.Printf("최종: %v (길이: %d, 용량: %d)\n", smallSlice, len(smallSlice), cap(smallSlice))
    
    fmt.Println("\n=== 용량 초과 시 내부 동작 시뮬레이션 ===")
    
    // 용량 초과 시 Go가 내부적으로 하는 일들을 시뮬레이션
    demonstrateCapacityGrowth()
}

func demonstrateCapacityGrowth() {
    fmt.Println("용량 초과 시 Go의 내부 동작:")
    fmt.Println("1. 더 큰 새로운 배열을 메모리에 할당")
    fmt.Println("2. 기존 요소들을 새 배열로 복사")
    fmt.Println("3. 새 요소를 추가")
    fmt.Println("4. 새 배열을 가리키는 슬라이스 반환")
    
    // 실제로 메모리 주소가 바뀌는지 확인
    slice := make([]int, 2, 2)
    slice[0] = 10
    slice[1] = 20
    
    fmt.Printf("\n초기 슬라이스 주소: %p\n", slice)
    fmt.Printf("초기 상태: %v (길이: %d, 용량: %d)\n", slice, len(slice), cap(slice))
    
    // 용량 초과 시 주소 변화 확인
    slice = append(slice, 30)
    fmt.Printf("append 후 주소: %p\n", slice)
    fmt.Printf("append 후: %v (길이: %d, 용량: %d)\n", slice, len(slice), cap(slice))
    
    if cap(slice) > 2 {
        fmt.Println("✅ 용량이 증가했고, 새로운 메모리 공간으로 이동했습니다!")
    }
}

// 용량 초과의 성능 영향을 보여주는 예제
func CapacityPerformanceExample() {
    fmt.Println("\n=== 용량 관리의 중요성 ===")
    
    // 나쁜 예: 용량을 고려하지 않은 경우
    fmt.Println("❌ 비효율적인 방법 (빈번한 메모리 재할당):")
    badSlice := make([]int, 0) // 용량을 지정하지 않음
    fmt.Printf("시작 용량: %d\n", cap(badSlice))
    
    reallocations := 0
    oldCap := cap(badSlice)
    
    for i := 0; i < 1000; i++ {
        badSlice = append(badSlice, i)
        if cap(badSlice) != oldCap {
            reallocations++
            oldCap = cap(badSlice)
        }
    }
    fmt.Printf("1000개 요소 추가 시 메모리 재할당 횟수: %d\n", reallocations)
    
    // 좋은 예: 예상 크기를 미리 지정
    fmt.Println("\n✅ 효율적인 방법 (용량 미리 할당):")
    goodSlice := make([]int, 0, 1000) // 예상 크기만큼 용량 미리 할당
    fmt.Printf("시작 용량: %d\n", cap(goodSlice))
    
    reallocations = 0
    oldCap = cap(goodSlice)
    
    for i := 0; i < 1000; i++ {
        goodSlice = append(goodSlice, i)
        if cap(goodSlice) != oldCap {
            reallocations++
            oldCap = cap(goodSlice)
        }
    }
    fmt.Printf("1000개 요소 추가 시 메모리 재할당 횟수: %d\n", reallocations)
    
    fmt.Println("\n💡 팁: 대략적인 크기를 알고 있다면 make([]T, 0, 예상크기)로 생성하세요!")
}