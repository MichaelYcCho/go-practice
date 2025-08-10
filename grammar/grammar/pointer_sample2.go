package grammar

import (
	"fmt"
	"unsafe"
)

// --- 심화 포인터 학습: 고급 개념과 실전 활용 ---

func PointersAdvanced() {
    fmt.Println("=== 포인터 심화 학습 ===")
    
    // --- 1. 함수 포인터 (Function Pointers) ---
    fmt.Println("\n--- 1. 함수 포인터 ---")
    demonstrateFunctionPointers()
    
    // --- 2. 메서드 포인터와 리시버 ---
    fmt.Println("\n--- 2. 메서드 포인터와 리시버 ---")
    demonstrateMethodPointers()
    
    // --- 3. 인터페이스와 포인터 ---
    fmt.Println("\n--- 3. 인터페이스와 포인터 ---")
    demonstrateInterfacePointers()
    
    // --- 4. 포인터 산술과 unsafe 패키지 ---
    fmt.Println("\n--- 4. unsafe 패키지와 저수준 메모리 조작 ---")
    demonstrateUnsafePointers()
    
    // --- 5. 메모리 레이아웃과 정렬 ---
    fmt.Println("\n--- 5. 메모리 레이아웃과 정렬 ---")
    demonstrateMemoryLayout()
}

// ---------------------- 함수 포인터 예제 ----------------------

// 함수 타입 정의
type MathOperation func(int, int) int
type StringProcessor func(string) string

func add(a, b int) int { return a + b }
func multiply(a, b int) int { return a * b }
func toUpper(s string) string { return fmt.Sprintf("UPPER: %s", s) }
func addPrefix(s string) string { return fmt.Sprintf("PREFIX_%s", s) }

func demonstrateFunctionPointers() {
    fmt.Println("함수 포인터로 동적 함수 선택:")
    
    // 함수 포인터 슬라이스
    operations := []MathOperation{add, multiply}
    operationNames := []string{"덧셈", "곱셈"}
    
    x, y := 5, 3
    for i, op := range operations {
        result := op(x, y)
        fmt.Printf("  %s: %d %s %d = %d\n", operationNames[i], x, 
                   map[int]string{0: "+", 1: "*"}[i], y, result)
    }
    
    // 함수를 매개변수로 전달
    fmt.Println("\n함수를 매개변수로 전달:")
    processStrings([]string{"hello", "world"}, toUpper)
    processStrings([]string{"test", "data"}, addPrefix)
    
    // 클로저와 함수 포인터
    fmt.Println("\n클로저와 함수 포인터:")
    multiplier := createMultiplier(10)
    fmt.Printf("  multiplier(5) = %d\n", multiplier(5))
}

func processStrings(strings []string, processor StringProcessor) {
    for _, s := range strings {
        fmt.Printf("  %s -> %s\n", s, processor(s))
    }
}

func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// ---------------------- 메서드 포인터 예제 ----------------------

// 계산기 구조체
type Calculator struct {
    value  int
    memory int
}

// 값 리시버 메서드
func (c Calculator) GetValue() int {
    return c.value
}

// 포인터 리시버 메서드 - 상태 변경
func (c *Calculator) Add(x int) {
    c.value += x
}

func (c *Calculator) Multiply(x int) {
    c.value *= x
}

func (c *Calculator) SaveToMemory() {
    c.memory = c.value
}

func (c *Calculator) LoadFromMemory() {
    c.value = c.memory
}

// 포인터 리시버 메서드 - 상태 초기화
func (c *Calculator) Reset() {
    c.value = 0
    c.memory = 0
}

func demonstrateMethodPointers() {
    calc := &Calculator{value: 10}
    fmt.Printf("초기값: %d\n", calc.GetValue())
    
    // 메서드 포인터를 변수에 저장
    addMethod := calc.Add
    multiplyMethod := calc.Multiply
    
    fmt.Println("메서드 포인터 사용:")
    addMethod(5)
    fmt.Printf("  Add(5) 후: %d\n", calc.GetValue())
    
    multiplyMethod(2)
    fmt.Printf("  Multiply(2) 후: %d\n", calc.GetValue())
    
    // 메서드 체이닝 패턴
    fmt.Println("\n메서드 체이닝 패턴:")
    calc.Reset()
    calc.Add2(10).Multiply2(3).SaveToMemory2()
    fmt.Printf("  체이닝 후 값: %d, 메모리: %d\n", calc.value, calc.memory)
}

// 메서드 체이닝을 위한 포인터 반환
func (c *Calculator) Add2(x int) *Calculator {
    c.value += x
    return c
}

func (c *Calculator) Multiply2(x int) *Calculator {
    c.value *= x
    return c
}

func (c *Calculator) SaveToMemory2() *Calculator {
    c.memory = c.value
    return c
}

// ---------------------- 인터페이스와 포인터 ----------------------

// 인터페이스 정의
type Drawable interface {
    Draw()
    Area() float64
}

type Resizable interface {
    Resize(factor float64)
}

// 도형 인터페이스 (Drawable + Resizable)
type Shape interface {
    Drawable
    Resizable
}

// 원 구조체
type Circle struct {
    radius float64
    x, y   float64
}

// 포인터 리시버로 인터페이스 구현
func (c *Circle) Draw() {
    fmt.Printf("  원 그리기: 중심(%.1f, %.1f), 반지름 %.1f\n", c.x, c.y, c.radius)
}

func (c *Circle) Area() float64 {
    return 3.14159 * c.radius * c.radius
}

func (c *Circle) Resize(factor float64) {
    c.radius *= factor
}

// 사각형 구조체
type Rectangle struct {
    width, height float64
    x, y         float64
}

func (r *Rectangle) Draw() {
    fmt.Printf("  사각형 그리기: 위치(%.1f, %.1f), 크기 %.1f x %.1f\n", 
               r.x, r.y, r.width, r.height)
}

func (r *Rectangle) Area() float64 {
    return r.width * r.height
}

func (r *Rectangle) Resize(factor float64) {
    r.width *= factor
    r.height *= factor
}

func demonstrateInterfacePointers() {
    // 인터페이스 슬라이스에 포인터 저장
    shapes := []Shape{
        &Circle{radius: 5, x: 0, y: 0},
        &Rectangle{width: 10, height: 20, x: 5, y: 5},
    }
    
    fmt.Println("도형들 그리기 및 크기 조절:")
    for i, shape := range shapes {
        fmt.Printf("도형 %d:\n", i+1)
        shape.Draw()
        fmt.Printf("  면적: %.2f\n", shape.Area())
        
        shape.Resize(1.5)
        fmt.Printf("  1.5배 확대 후 면적: %.2f\n", shape.Area())
    }
    
    // 타입 단언과 포인터
    fmt.Println("\n타입 단언과 포인터:")
    for i, shape := range shapes {
        switch s := shape.(type) {
        case *Circle:
            fmt.Printf("도형 %d는 반지름 %.1f인 원입니다\n", i+1, s.radius)
        case *Rectangle:
            fmt.Printf("도형 %d는 %.1f x %.1f인 사각형입니다\n", i+1, s.width, s.height)
        }
    }
}

// ---------------------- unsafe 패키지 예제 ----------------------

func demonstrateUnsafePointers() {
    fmt.Println("⚠️ 주의: unsafe 패키지는 저수준 프로그래밍에서만 사용")
    
    // 기본 타입의 메모리 크기
    var i int = 42
    var f float64 = 3.14
    var s string = "hello"
    
    fmt.Printf("int 크기: %d 바이트\n", unsafe.Sizeof(i))
    fmt.Printf("float64 크기: %d 바이트\n", unsafe.Sizeof(f))
    fmt.Printf("string 크기: %d 바이트\n", unsafe.Sizeof(s))
    
    // 구조체의 필드 오프셋
    type Person struct {
        Name string
        Age  int
        ID   int64
    }
    
    var p Person
    fmt.Printf("\nPerson 구조체 분석:\n")
    fmt.Printf("전체 크기: %d 바이트\n", unsafe.Sizeof(p))
    fmt.Printf("Name 오프셋: %d 바이트\n", unsafe.Offsetof(p.Name))
    fmt.Printf("Age 오프셋: %d 바이트\n", unsafe.Offsetof(p.Age))
    fmt.Printf("ID 오프셋: %d 바이트\n", unsafe.Offsetof(p.ID))
    
    // 포인터 변환 (매우 위험!)
    fmt.Println("\n포인터 변환 예제 (교육용):")
    x := int64(1000)
    ptr := unsafe.Pointer(&x)
    
    // int64 포인터를 []byte로 해석
    byteSlice := (*[8]byte)(ptr)
    fmt.Printf("int64 값 %d의 바이트 표현: %v\n", x, byteSlice[:])
}

// ---------------------- 메모리 레이아웃 예제 ----------------------

func demonstrateMemoryLayout() {
    // 구조체 패딩과 정렬
    type BadAlignment struct {
        a bool    // 1바이트
        b int64   // 8바이트
        c bool    // 1바이트
        d int32   // 4바이트
    }
    
    type GoodAlignment struct {
        b int64   // 8바이트
        d int32   // 4바이트
        a bool    // 1바이트
        c bool    // 1바이트
        // 2바이트 패딩
    }
    
    var bad BadAlignment
    var good GoodAlignment
    
    fmt.Printf("잘못된 정렬 구조체 크기: %d 바이트\n", unsafe.Sizeof(bad))
    fmt.Printf("올바른 정렬 구조체 크기: %d 바이트\n", unsafe.Sizeof(good))
    
    // 슬라이스의 내부 구조
    fmt.Println("\n슬라이스 내부 구조:")
    slice := []int{1, 2, 3, 4, 5}
    
    fmt.Printf("슬라이스 헤더 크기: %d 바이트\n", unsafe.Sizeof(slice))
    fmt.Printf("길이: %d, 용량: %d\n", len(slice), cap(slice))
    fmt.Printf("데이터 포인터: %p\n", &slice[0])
    
    // 문자열의 내부 구조
    fmt.Println("\n문자열 내부 구조:")
    str := "Hello, Go!"
    fmt.Printf("문자열 헤더 크기: %d 바이트\n", unsafe.Sizeof(str))
    fmt.Printf("문자열 길이: %d\n", len(str))
}

// ---------------------- 고급 포인터 패턴 ----------------------

// 옵션 패턴을 위한 포인터 활용
type ServerConfig struct {
    Host    string
    Port    int
    Timeout *int  // nil이면 기본값 사용
    Debug   *bool // nil이면 기본값 사용
}

func NewServerConfig(host string, port int) *ServerConfig {
    return &ServerConfig{
        Host: host,
        Port: port,
    }
}

func (c *ServerConfig) WithTimeout(timeout int) *ServerConfig {
    c.Timeout = &timeout
    return c
}

func (c *ServerConfig) WithDebug(debug bool) *ServerConfig {
    c.Debug = &debug
    return c
}

func (c *ServerConfig) GetTimeout() int {
    if c.Timeout != nil {
        return *c.Timeout
    }
    return 30 // 기본값
}

func (c *ServerConfig) IsDebug() bool {
    if c.Debug != nil {
        return *c.Debug
    }
    return false // 기본값
}

func AdvancedPointerPatterns() {
    fmt.Println("\n=== 고급 포인터 패턴 ===")
    
    // 빌더 패턴
    fmt.Println("빌더 패턴:")
    config := NewServerConfig("localhost", 8080).
        WithTimeout(60).
        WithDebug(true)
    
    fmt.Printf("서버 설정: %s:%d, 타임아웃: %d초, 디버그: %t\n",
        config.Host, config.Port, config.GetTimeout(), config.IsDebug())
    
    // nil 포인터를 활용한 옵셔널 필드
    fmt.Println("\n옵셔널 필드:")
    config2 := NewServerConfig("example.com", 443)
    fmt.Printf("기본 설정: 타임아웃: %d초, 디버그: %t\n",
        config2.GetTimeout(), config2.IsDebug())
}