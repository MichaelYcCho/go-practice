package grammar

import "fmt"

// User라는 이름의 구조체(struct)를 정의합니다.
// 구조체는 여러 개의 필드(변수)를 묶어서 새로운 타입을 만듭니다.
// 각 필드는 이름과 타입을 가집니다.
type User struct {
    Name  string
    Age   int
    Email string
}

func StructSample() {
    // --- 1. 구조체 변수 선언 및 초기화 ---

    // 1-1. 필드 이름을 명시하여 초기화합니다.
    // 필드 순서와 상관없이 초기화할 수 있어 가독성이 좋습니다.
    user1 := User{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }
    fmt.Println("--- 1. 구조체 변수 초기화 ---")
    fmt.Printf("user1: %+v\n", user1) // %+v는 필드 이름과 값을 모두 출력합니다.

    // 1-2. 필드 이름을 생략하고 순서대로 초기화합니다.
    // 필드 순서를 정확히 알아야 하기 때문에 일반적으로 잘 사용하지 않습니다.
    user2 := User{"Bob", 25, "bob@example.com"}
    fmt.Printf("user2: %+v\n", user2)
    
    // --- 2. 구조체 필드 접근 및 수정 ---
    
    // 마침표(.) 연산자를 사용하여 구조체의 필드에 접근하고 값을 수정할 수 있습니다.
    fmt.Println("\n--- 2. 필드 접근 및 수정 ---")
    fmt.Printf("user1의 이름: %s\n", user1.Name) // user1의 Name 필드에 접근
    
    user1.Age = 31 // user1의 Age 필드 값을 수정
    fmt.Printf("user1의 수정된 나이: %d\n", user1.Age)

    // --- 3. 구조체에 메서드 호출 ---
    
    // 이제 `User` 구조체에 속한 메서드를 호출해 보겠습니다.
    fmt.Println("\n--- 3. 메서드 호출 ---")
    
    // user1에 속한 `Introduce` 메서드를 호출합니다.
    // 메서드는 구조체 변수를 통해 호출됩니다.
    user1.Introduce()

    // user1에 속한 `UpdateAge` 메서드를 호출합니다.
    // 이 메서드는 `포인터 리시버`를 사용하여 원본 구조체의 값을 직접 수정합니다.
    user1.UpdateAge(35)
    fmt.Printf("UpdateAge 메서드 호출 후 user1의 나이: %d\n", user1.Age)
}

// ---------------------- 메서드 정의 ----------------------

// Introduce 메서드: `User` 구조체에 속한 함수입니다.
// func (리시버) 함수이름(매개변수) 반환타입
// `(u User)`는 이 메서드가 `User` 타입에 속한다는 것을 의미하는 `리시버(receiver)`입니다.
// u는 메서드 내부에서 `User` 구조체 인스턴스를 가리키는 변수명입니다.
func (u User) Introduce() {
    fmt.Printf("안녕하세요, 저는 %s이고, %d살입니다.\n", u.Name, u.Age)
}

// UpdateAge 메서드: `User` 구조체의 나이를 업데이트합니다.
// `(u *User)`와 같이 리시버에 포인터(*)를 사용하면, 원본 구조체의 값을 직접 수정할 수 있습니다.
// 만약 포인터가 아니었다면(u User), 복사된 구조체의 값만 수정되고 원본은 바뀌지 않습니다.
func (u *User) UpdateAge(newAge int) {
    u.Age = newAge
}