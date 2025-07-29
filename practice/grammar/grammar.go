package grammar

import "fmt" // fmt 패키지는 입출력(format) 관련 함수들을 제공합니다.

func For() { // main 함수는 프로그램의 시작점입니다.

    // --- 1. 기본적인 for 루프 (C/Java 스타일) ---
    // 초기화(initialization); 조건(condition); 후처리(post-statement)
    // 조건이 참(true)인 동안 반복됩니다.
    fmt.Println("--- 1. 기본적인 for 루프 (1부터 5까지 출력) ---")
    for i := 1; i <= 5; i++ { // i를 1로 초기화하고, i가 5보다 작거나 같을 때까지 i를 1씩 증가시키며 반복
        fmt.Printf("현재 숫자: %d\n", i)
    }
    fmt.Println("--------------------------------------------------")

    // --- 2. 조건만 있는 for 루프 (While 루프처럼 사용) ---
    // 초기화와 후처리 구문을 생략하고 조건만으로 반복할 수 있습니다.
    // 이는 다른 언어의 'while' 루프와 유사합니다.
    fmt.Println("\n--- 2. 조건만 있는 for 루프 (While처럼 사용) ---")
    j := 0 // 반복문 밖에서 변수 j를 선언 및 초기화
    for j < 3 { // j가 3보다 작을 때까지 반복
        fmt.Printf("현재 j: %d\n", j)
        j++ // 반복문 내부에서 j 값을 증가시킴
    }
    fmt.Println("--------------------------------------------------")

    // --- 3. 무한 루프 (Infinite Loop) ---
    // 조건 구문을 생략하면 무한히 반복됩니다.
    // 보통 'break' 문과 함께 사용하여 특정 조건에서 루프를 종료합니다.
    fmt.Println("\n--- 3. 무한 루프 (break 사용) ---")
    k := 0
    for { // 조건이 없으므로 무한 반복
        fmt.Printf("현재 k: %d\n", k)
        k++
        if k >= 3 { // k가 3 이상이 되면 루프를 종료
            break // 루프를 즉시 종료하고 다음 코드로 넘어갑니다.
        }
    }
    fmt.Println("무한 루프 종료!")
    fmt.Println("--------------------------------------------------")

    // --- 4. for-range 루프 (컬렉션 순회) ---
    // 배열(array), 슬라이스(slice), 문자열(string), 맵(map), 채널(channel)과 같은 컬렉션을 순회할 때 사용합니다.
    // 각 반복에서 인덱스와 값을 반환합니다.
    fmt.Println("\n--- 4. for-range 루프 (슬라이스 순회) ---")
    numbers := []int{10, 20, 30, 40, 50} // 정수형 슬라이스 선언 및 초기화

    for index, value := range numbers { // numbers 슬라이스의 각 요소에 대해 반복. index와 value를 반환
        fmt.Printf("인덱스: %d, 값: %d\n", index, value)
    }
    fmt.Println("--------------------------------------------------")

    // for-range 루프에서 인덱스만 필요하거나 값만 필요할 때
    fmt.Println("\n--- 4-1. for-range 루프 (값만 필요할 때) ---")
    for _, value := range numbers { // 인덱스가 필요 없으면 '_' (블랭크 식별자)를 사용
        fmt.Printf("값: %d\n", value)
    }
    fmt.Println("--------------------------------------------------")

    fmt.Println("\n--- 4-2. for-range 루프 (인덱스만 필요할 때) ---")
    for index := range numbers { // 값만 필요 없으면 두 번째 변수를 생략
        fmt.Printf("인덱스: %d\n", index)
    }
    fmt.Println("--------------------------------------------------")

    // --- 5. continue 문 ---
    // 'continue' 문은 현재 반복을 건너뛰고 다음 반복으로 넘어갑니다.
    fmt.Println("\n--- 5. continue 문 (짝수만 출력) ---")
    for i := 1; i <= 5; i++ {
        if i%2 != 0 { // i가 홀수이면
            continue // 현재 반복을 건너뛰고 다음 i 값으로 넘어감
        }
        fmt.Printf("짝수: %d\n", i)
    }
    fmt.Println("--------------------------------------------------")

    // --- 6. 중첩 루프 (Nested Loops) ---
    // 루프 안에 또 다른 루프를 넣을 수 있습니다.
    // 구구단 예시
    fmt.Println("\n--- 6. 중첩 루프 (구구단 2단~4단) ---")
    for dan := 2; dan <= 4; dan++ { // 바깥 루프: 단 (2단부터 4단까지)
        fmt.Printf("\n--- %d단 ---\n", dan)
        for num := 1; num <= 9; num++ { // 안쪽 루프: 곱하는 수 (1부터 9까지)
            fmt.Printf("%d * %d = %d\n", dan, num, dan*num)
        }
    }
    fmt.Println("--------------------------------------------------")
}