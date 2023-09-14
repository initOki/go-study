package main

import (
	"fmt"
	"math"
)

func main() {
	// 대문자로 시작하는 이름은 export됨
	// package를 import 할 때 해당 패키지에 export 된 이름들만 참조할 수 있음
	// "fmt.Println(math.pi)"는 에러를 발생시킴
	fmt.Println(math.Pi)
}
