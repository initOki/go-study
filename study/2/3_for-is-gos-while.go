package main

import "fmt"

func main() {
	sum := 1
	// ;을 생략 할 수 있다는 점에서 C에서의 while문과 비슷하다.
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
