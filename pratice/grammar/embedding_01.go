package grammar

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // 타입의 항목을 포함하지만 항목의 이름이 지정되어있지 않은것 = 즉 임베딩 한 것이다
	Reports  []Employee
}

func (m Manager) FindNewEmployee() []Employee {
	// do Something
	return m.Reports
}

// 임베딩된 항목의 선언된 모든 항목이나 메서드는 승격되어 구조체를 포함하고 바로 실행이 가능하다
func EmbeddingStart01() {
	m := Manager{
		Employee: Employee{
			Name: "John",
			ID:   "123",
		},
		Reports: []Employee{},
	}

	fmt.Println(m.ID)
	fmt.Println(m.Description())
}

// 임베딩 케이스 2

type Inner01 struct {
	X int
}

type Outer01 struct {
	Inner01
	X int
}

func EmbeddingStart02() {
	o := Outer01{
		Inner01: Inner01{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)         // 20
	fmt.Println(o.Inner01.X) //10 명시적으로 임베딩항목을 정의하여 내부값에 접근할 수 있다.
}
