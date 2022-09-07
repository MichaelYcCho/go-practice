package main

import "fmt"

func main() {
	c := make(chan int)

	//send (채널 c를 전달)
	go sendingChennel(c)

	//receive (채널 c를 전달)
	receiveChennel(c)

	fmt.Println("about to exit")
}

func sendingChennel(c chan<- int) {
	c <- 42
}

func receiveChennel(c <-chan int) {
	fmt.Println(<-c)
}
