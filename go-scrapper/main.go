package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Rice", "ramen"}
	michael := person{name: "michael", age: 10, favFood: favFood}
	fmt.Println(michael.name)

}
