package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 5

	if a > b {
		fmt.Println("Si,", a, "es mayor que", b)
	} else {
		fmt.Println("No,", a, "no es mayor que", b)
	}

	a = 5
	if a = a - 1; a >= b {
		fmt.Println(a, "es mayor o igual a", b)
	} else {
		fmt.Println("Si,", b, "es mayor que", a)
	}

	c := 5
	if c == 1 {
		fmt.Println("C es igual a 1")
	} else if c == 2 {
		fmt.Println("C es igual a 2")
	} else {
		fmt.Println("No se cuanto es C")
	}
}
