package grammar

import (
	"fmt"
	"reflect"
)

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		// DeepEqual은 슬라이스를 포함한 모든것에 대해 비교가 가능하다
		return reflect.DeepEqual(me, me2)
	}
	return false

}
