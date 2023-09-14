package main

import "fmt"

func main() {
	// v type은 string
	v := "42"
	// i type은 float64
	i := 3.142
	// g type은 complex128
	g := 0.867 + 0.5i

	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("g is of type %T\n", g)
}
