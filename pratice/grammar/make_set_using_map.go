package grammar

import "fmt"

// Go에서는 공식적으로 set을 지원하지 않으므로 맵을 활용해야한다.
// 맵은 중복되는 키를 허용하지 않는것을 활용하는 것이다.

// bool을 활용한 경우
func MakeSetUsingMap01() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])   // 5라는 key가 존재하므로 true
	fmt.Println(intSet[500]) // 500이라는 key가 존재하지 않으므로 false
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

// stuct{}를 활용하면 0바이트를 할당해서 1바이트를할당하는 bool보다 메모리 사용량이 효율적이다. 다만 극적인 차이는 없다.
func MakeSetUsingMap02() {
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = struct{}{}
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}

	fmt.Println(len(vals), len(intSet), intSet)
}
