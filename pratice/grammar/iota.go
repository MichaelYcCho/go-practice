package grammar

import "fmt"

func Iota_base() {
	const (
		a = iota
		b = iota
		c = iota
	)

	const (
		d = iota
		e
		f
	)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}

func Iota_bit() {

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Println()

	const (
		_ = iota
		// kb = 24
		kb_1 = 1 << (iota * 10)
		mb_1 = 1 << (iota * 10)
		gb_1 = 1 << (iota * 10)
	)
	fmt.Printf("%d\t\t%b\n", kb_1, kb_1)
	fmt.Printf("%d\t\t%b\n", mb_1, mb_1)
	fmt.Printf("%d\t%b\n", gb_1, gb_1)

}
