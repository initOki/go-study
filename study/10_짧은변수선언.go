package main

import "fmt"

func main() {
	var i, j int = 23, 21
	// 함수 내에서 사용되는 :=라는 짧은 변수 선언은 암시적으로 type으로 var 선언처럼 사용될 수 있음
	// 단 함수 밖에서 선언된 경우에는 :=를 사용할 수 없음
	k := 35
	c, python, java := true, false, "Hi!"
	fmt.Println(i, j, k, c, python, java)
}
