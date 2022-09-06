package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 42
	fmt.Println(<-c)

}

/*
동시에 실행될 수 없다면 받는 대상이 데이터를 빼낼수 있을때까지 송수신은 차단된다.
위의 경우 main함수에 진입해서 채널을 생성한 뒤로 채널이 수신하는 작업이 없기때문에 멈춰버린다.

*/
