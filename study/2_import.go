package main

// ()를 사용해서 import를 그룹지을 수 있음
// 다른 예시로
// import "fmt"
// import "math"
// 와 같이 작성 할 수 있음
// 하지만 나눠서 사용하는 것 보다는 괄호를 이용해서 그룹지어서 사용하는 것을 추천함
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
