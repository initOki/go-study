package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// float64로 type 변환
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// unit로 type 변환
	var z uint = uint(f)

	fmt.Println(x, y, z)
}
