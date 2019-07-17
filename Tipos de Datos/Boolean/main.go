package main

import (
	"fmt"
)

func main() {
	var resultado bool

	resultado = 5 < 6
	fmt.Println("5 < 6 = ", resultado)

	resultado = (5 > 6) && (4 < 3)
	fmt.Println("(5 > 6) && (4 < 3) = ", resultado)

	resultado = (5 > 6) || (4 > 3)
	fmt.Println("(5 > 6) || (4 > 3) = ", resultado)

	resultado = (5 > 6) || (4 > 3)
	fmt.Println("!((5 > 6) || (4 > 3)) = ", !resultado)
}
