package grammar

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker02(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// %w 다른 오류의 형식으로 지정된 문자열과 원본 오류를 출력한다.
		// 관례상 %w 뒤에 error를 넣는다
		return fmt.Errorf("in fileChecker %w", err)
	}
	f.Close()
	return nil
}

func ErrorTest04_01() {
	err := fileChecker02("test.txt")
	if err != nil {
		// Is는 해당 함수와 확인될 오류와 대응된느 인스턴스를 받아와 오류가 있다면 true를 반환한다.
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file doesn't exist")
		}
	}
}
