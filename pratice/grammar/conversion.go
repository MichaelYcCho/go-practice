package grammar

// 다른곳에선 casting이라고하지만 Go에선 conversion(변환) 이라고 한다.

var a int

type hotdog int

var b hotdog

func TypeChanging() {
	a = 42
	b = 43

	// a = b // error
	// b = a // error

	// type changing
	a = int(b)
	b = hotdog(a)
}
