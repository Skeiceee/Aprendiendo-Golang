package main

import (
	"fmt"
)

func main(){
	var j []int
	fmt.Println("Slice j:", j)
	
	x := []int{1, 2, 3}
	fmt.Println("Slice x:", x)

	y := make([]int, 5)
	fmt.Println("Slice y:", y)
	fmt.Println("Longitud de y:", len(y))
	fmt.Println("Capacidad de y:", cap(y))

	k := make([]int, 5, 10)
	fmt.Println("Slice k:", k)
	fmt.Println("Longitud de k:", len(k))
	fmt.Println("Capacidad de k:", cap(k))

	var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Array ar:", arr)

	var a, b []int
	fmt.Println("Slice a:", a)
	fmt.Println("Slice b:", b)	

	a = arr[2:5]
	b = arr[3:5]
	fmt.Println("Slice a con contenido de arr:", a)
	fmt.Println("Slice b con contenido de arr:", b)

	b[0] = 25
	fmt.Println("Asignamos b[0] = 25")
	
	fmt.Println("Slice b:", b)
	fmt.Println("Slice a:", a)
	fmt.Println("Array arr:", arr)
	
	fmt.Println("Capacidad de b:", cap(b))
	fmt.Println("Capacidad de a:", cap(a))
}