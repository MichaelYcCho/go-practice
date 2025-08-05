package grammar

import "fmt"

// Speaker라는 이름의 인터페이스를 정의합니다.
// 이 인터페이스는 Speak()라는 메서드(반환값과 매개변수가 없는)를 가집니다.
// 어떤 타입이든 Speak() 메서드를 구현하면 Speaker 인터페이스로 취급될 수 있습니다.
type Speaker interface {
    Speak()
}

// Person 구조체는 Name 필드를 가집니다.
type Person struct {
    Name string
}

// Animal 구조체는 Species(종) 필드를 가집니다.
type Animal struct {
    Species string
}

// Person 구조체를 위한 Speak() 메서드를 구현합니다.
// Person은 이제 Speaker 인터페이스를 '구현'한 것입니다.
func (p Person) Speak() {
    fmt.Printf("%s가 말합니다: 안녕하세요!\n", p.Name)
}

// Animal 구조체를 위한 Speak() 메서드를 구현합니다.
// Animal도 이제 Speaker 인터페이스를 '구현'한 것입니다.
func (a Animal) Speak() {
    fmt.Printf("%s가 말합니다: 멍멍!\n", a.Species)
}

// --- 인터페이스를 활용하는 함수 ---

// Introduce 함수는 Speaker 인터페이스 타입의 인자를 받습니다.
// 따라서 Person 타입이든 Animal 타입이든 Speak() 메서드가 구현된
// 모든 타입을 인자로 받을 수 있습니다.
func Introduce(s Speaker) {
    s.Speak() // s의 실제 타입에 따라 다른 Speak() 메서드가 호출됩니다.
}

func main() {
    // --- 1. 인터페이스를 구현한 구조체 인스턴스 생성 ---

    // Person 구조체 인스턴스를 생성합니다.
    p := Person{Name: "철수"}

    // Animal 구조체 인스턴스를 생성합니다.
    a := Animal{Species: "강아지"}

    // --- 2. 인터페이스를 활용한 다형성 ---
    
    fmt.Println("--- 2. Introduce 함수를 통한 다형성 ---")

    // Introduce 함수에 Person 타입의 인자를 전달합니다.
    // Person은 Speak()를 구현했으므로, Speaker 인터페이스로 취급됩니다.
    Introduce(p)

    // Introduce 함수에 Animal 타입의 인자를 전달합니다.
    // Animal도 Speak()를 구현했으므로, Speaker 인터페이스로 취급됩니다.
    Introduce(a)

    // --- 3. 인터페이스 변수에 값 할당 ---
    
    fmt.Println("\n--- 3. 인터페이스 변수에 값 할당 ---")

    // s 변수는 Speaker 인터페이스 타입입니다.
    // Speak() 메서드가 구현된 모든 타입을 할당할 수 있습니다.
    var s Speaker
    
    // s에 Person 구조체 인스턴스를 할당합니다.
    s = p
    s.Speak() // "철수가 말합니다: 안녕하세요!" 출력

    // s에 Animal 구조체 인스턴스를 할당합니다.
    s = a
    s.Speak() // "강아지가 말합니다: 멍멍!" 출력
    
    // --- 4. 타입 단언(Type Assertion) ---
    
    fmt.Println("\n--- 4. 타입 단언 ---")
    
    // 인터페이스 변수에 저장된 값의 실제 타입을 확인하고 싶을 때 사용합니다.
    // a.(Animal)은 s의 실제 타입이 Animal이라고 '단언'하는 것입니다.
    // 만약 실제 타입이 Animal이 아니면 런타임 패닉(panic)이 발생합니다.
    s = a
    animal, ok := s.(Animal)
    if ok {
        fmt.Printf("s의 실제 타입은 Animal입니다. 종: %s\n", animal.Species)
    }
    
    // Person 타입에 대해 단언하면 실패합니다.
    _, ok = s.(Person)
    if !ok {
        fmt.Printf("s의 실제 타입은 Person이 아닙니다.\n")
    }
}