package grammar

// CreateCounter 함수는 클로저(익명 함수)를 반환합니다.
// 반환되는 함수는 int 타입의 값을 반환하는 함수입니다.
func CreateCounter() func() int {
    // count 변수는 이 함수의 스코프 내에 있습니다.
    // 하지만 아래의 익명 함수가 이 변수를 "기억"하고 있기 때문에,
    // 이 함수의 실행이 끝나도 count 변수는 사라지지 않습니다.
    count := 0

    // 이 익명 함수가 클로저입니다.
    // 이 함수는 외부 스코프에 있는 'count' 변수에 접근하여 값을 변경하고 반환합니다.
    return func() int {
        count++ // 이 익명 함수가 호출될 때마다 count 변수의 값을 1씩 증가시킵니다.
        return count // 증가된 값을 반환합니다.
    }
}