package main

import "fmt"

// 함수는 0개 또는 그 이상의 인자를 받을 수 있다.
// 변수 이름 뒤에는 type을 적어준다.
// x type, y type을 받아서 int type을 반환한다.
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 13))
}
