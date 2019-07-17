package main

import (
	"fmt"
)

func main() {
	//	Float
	var f32 float32
	var f64 float64
	var c64 complex64
	var c184 complex128

	fmt.Println("f32, f64, c64, c184 = ", f32, f64, c64, c184)

	f32 = .156
	f64 = .156
	fmt.Println("f32 * 58.565656 = ", f32*58.565656)
	fmt.Println("f64 * 58.565656 = ", f64*58.565656)

	c64 = complex(5, 6)
	fmt.Println("c64 = ", c64)
	fmt.Println("c64 * 85458.65", c64*85458.65)
	c184 = complex(5, 6)
	fmt.Println("c184 = ", c184)
	fmt.Println("c184 * 85458.65= ", c184*85458.65)

	fmt.Println("Real - c184 * 85458.65= ", real(c184*85458.65))
	fmt.Println("Imaginaria - c184 * 85458.65= ", imag(c184*85458.65))
}
