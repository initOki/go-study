package main

import (
	"fmt"
	"math/cmplx"
)

// Go의 기본 타입은 아래와 같다
//bool

//string

//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr

//byte // uint8의 별칭

//rune // int32의 별칭
// 유니코드에서 code point를 의미합니다.

//float32 float64

//complex64 complex128

// 변수 선언들은 import 문과 마찬가지로 조각으로 쪼갤 수 있음
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
