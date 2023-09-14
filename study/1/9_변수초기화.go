package main

import "fmt"

// 변수 선언은 한 변수당 하나의 초기값을 가질 수 있음
// 만약에 초기값이 있을 경우 type은 생략될 수 있음
// 이런 경우에는 변수는 초기값의 type을 가짐
var i, j int = 123, 21

func main() {
	// bool, bool, string
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
