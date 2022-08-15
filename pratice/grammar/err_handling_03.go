package grammar

import "fmt"

// 자신만의 오류타입을 사용할때는 초기화되지 않은 인스턴스를 반환하면 안된다.

func GenerateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func GenerateErrorFix_01(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	// 첫번째 방법은 명시적으로 error 값을 nil로 반환하는 것이다.
	return nil
}

func GenerateErrorFix_02(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	// 첫번째 방법은 명시적으로 error 값을 nil로 반환하는 것이다.
	return genErr
}

func ErrorTest03() {
	// error는 인터페이스다. 따라서 인터페이스가 nil로 간주되려면 반드시 기본 타입과 값이 nil이어야한다.
	err := GenerateError(true)
	fmt.Println(err != nil) // true
	err = GenerateError(false)
	fmt.Println(err != nil) // true

}
