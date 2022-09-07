package main

import "fmt"

func main() {
	c := make(chan int)

	//send (채널 c를 전달)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	//receive (채널 c를 전달)
	receiveChennel(c)

	fmt.Println("about to exit")
}

func sendingChennel(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receiveChennel(c <-chan int) {
	for v := range c {
		fmt.Println(v)

	}
}
