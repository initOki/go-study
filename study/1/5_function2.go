package main

import "fmt"

// 두 개 이상의 연속된 이름의 동일한 type일 경우에는 마지막 type을 제외한 나머지에 type은 생략 할 수 있음
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add2(1, 13))
}
