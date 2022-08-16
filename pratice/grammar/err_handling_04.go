package grammar

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// %w 다른 오류의 형식으로 지정된 문자열과 원본 오류를 출력한다.
		// 관례상 %w 뒤에 error를 넣는다
		return fmt.Errorf("in fileChecker %w", err)
	}
	f.Close()
	return nil
}

func ErrorTest04() {
	err := fileChecker("test.txt")
	if err != nil {
		fmt.Println("에러에요", err)
		// 래핑된 오류가 있으면 이를 반환하고 없으면 nil을 반환한다.
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println("래핑된 에러에요", wrappedErr)
		}
	}
}
