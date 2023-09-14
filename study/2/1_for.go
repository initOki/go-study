package main

import "fmt"

func main() {
	sum := 0
	// javascript와 다르게 소괄호를 사용하지 않는다.
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}
