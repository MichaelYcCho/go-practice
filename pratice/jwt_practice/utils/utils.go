package utils

import (
	"crypto/rand"
	"fmt"
)

// byte를 읽어와서 16진수 %X로 반환
func RandString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	str := fmt.Sprintf("%X", b)

	return str
}
