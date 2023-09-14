package main

import "fmt"

// 상수는 const 키워드를 사용하여 선언한다.
// 상수는 character , string , boolean , int 타입 중의 하나가 될 수 있다.
// 상수는 := 를 사용하여 선언할 수 없다.
const Pi = 3.14

func main() {
	const World = "안녕"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
