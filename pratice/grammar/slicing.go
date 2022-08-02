package grammar

import "fmt"

func Slicing() {

	// 길이가 0이고 수용력이 5인 슬라이스를 만든다.
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	fmt.Println("x", x)
	y := x[:2]
	z := x[2:]

	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
}

/*
기대값
x [1 2 3 4 60]
y [1 2 30 40 50]
z [3 4 70]

실제 출력값
5 5 3
x [1 2 30 40 70]
y [1 2 30 40 70]
z [30 40 70]

why? 다른 슬라이스로부터 슬라이스를 취할때마다
하위슬라이스의 수용력 =  원븐 슬라이스의 수용력 - 하위슬라이스  시작 오프셋만큼 뺀 값이 설정되기 때문
즉 원본 슬라이스의 사용되지 않은 모든 수용력은 만들어진 모든 하위 슬라이스에 공유가 됨
x에서 y 슬라이스를 만들때 길이는 2지만 수용력은 x와 동일한 4가 된다.
수용력이 4이므로 y에 끝에 append()를 하면 x의 세번째 위치에 요소가 추가된다
*/

func SlicingSoulution() {

	// 길이가 0이고 수용력이 5인 슬라이스를 만든다.
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)

	// 완전한 슬라이스 연산은 하위 슬라이스를 위해 가용한 부모슬라이스의 수용력의 마지막 위치를 지정하는 세번째 인자를 가진다.
	// 하위 슬라이스의 수용력을 계산하기 위해서는 세번쨰 인자에서 시작 오프셋을 빼면된다.
	//y와 z의 수용력은 둘다 2이다.
	y := x[:2:2]
	z := x[2:4:4]

	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
}

/*

5 2 2
x [1 2 3 4 60]
y [1 2 30 40 50]
z [3 4 70]

*/
