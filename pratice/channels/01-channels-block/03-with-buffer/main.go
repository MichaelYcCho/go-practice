package main

import "fmt"

/*
	버퍼 채널은 채널안에 값이 수신할 주체가 없다해도 머물수 있게해준다.
	위의 경우는 버퍼채널에 값 하나를 담을수 있게된다.

*/

func UnSuccessBuffer() {
	c := make(chan int, 1)

	c <- 42
	c <- 43

	fmt.Println(<-c)

}

func main() {

	//SuccessBuffer()
	UnSuccessBuffer()

}
