package main

import "fmt"

func main() {
	sum := 1
	// for문에서 초기화와 후처리를 생략할 수 있다.
	// 초기화와 후처리는 생략할 수 있지만 조건은 생략할 수 없다.
	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)
}
