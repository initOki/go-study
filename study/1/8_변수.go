package main

import "fmt"

// 변수에 대한 목록을 선언
// 마지막은 type으로 끝난다.
// var 문은 package 또는 function 안에서 사용할 수 있다.

// package에서 사용
var c, python, java bool

func main() {
	// function에서 사용
	// 초기화 하지 않은 변수는 0의 value를 갖음
	var i int
	fmt.Println(i, c, python, java)
}
