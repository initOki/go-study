package main

import "fmt"

// 인자가 없는 return 문은 이름이 주언진 반환 값을 반환한다.
// 이것을 "naked return"이라고 한다.
// * 이런 방식은 긴 함수에서 사용할 경우 가독성을 떨어뜨릴 수 있기때문에 주의해야한다.
func split(sum int) (x, y int) {
	x = sum * 4 / 6
	y = sum - x
	// naked return
	return
}

func main() {
	fmt.Println(split(14))
}
