package main

import "fmt"

// 하나의 함수는 여러개의 결과를 반환 할 수 있음
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("안녕", "Go~!")
	fmt.Println(a, b)
}
