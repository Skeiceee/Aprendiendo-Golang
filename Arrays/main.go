package main

import (
	"fmt"
)

func main() {
	var x [3]int
	fmt.Println(x)

	var k [3][2]float64
	fmt.Println(k)

	x[0] = 25
	fmt.Println(x[0])

	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	j := [...]int{1, 2, 3}
	fmt.Println(j)

	i := [...]int{
		1,
		2,
		3,
		4,
		5,
		6,
	}
	fmt.Println(i)

	f := [...]float64{1.2, 3.5, 2.5}
	fmt.Println(f)
	fmt.Println(len(f))

	fmt.Println(f[len(f)-1])

	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 2, 3}
	//d := [4]int{1, 2, 3, 4}

	fmt.Println(a == b)
	fmt.Println(b == c)
	//fmt.Println(c == d) (No se puede)
}
