package grammar

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
)

// 패키지 이름 재정의

func PackageTest01() *rand.Rand {

	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic(err)
	}
	r := rand.New(rand.NewSource(int64(b[0])))

	fmt.Println(&r)
	return r
}
