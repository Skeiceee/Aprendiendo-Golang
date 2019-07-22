package main

import (
	"fmt"
)

func main(){
	x := make([]byte, 4, 10)
	fmt.Println(x)

	x = []byte{'H', 'O', 'L', 'A'}
	fmt.Println(x)

	fmt.Printf("Slice x: %q \n", x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("Slice x[%d]: %q \n", i, x[i])
	}

	x = append(x, ' ') // Capacidad de 8
	fmt.Println(x, cap(x))

	x = append(x, 'M', 'U', 'N', 'D', 'O')
	fmt.Printf("Slice x: %q \n", x)
	fmt.Println(cap(x))

	var y []int
	for i := 0; i <= 15; i++ {
		y = append(y, i)
		fmt.Printf("Slice y[%d]: %d\t\t", i, i)
		fmt.Printf("Len y: %d\tCap y: %d\tElementos: %d\n", len(y), cap(y), i+1)
	}
	fmt.Println(y)
}