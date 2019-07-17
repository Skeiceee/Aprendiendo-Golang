package main

import (
	"fmt"
)

func main() {
	var dia int

	fmt.Print("Digite el numero del dia de la semana: ")
	fmt.Scanln(&dia)

	if dia == 1 {
		fmt.Println("El dia es Lunes.")
	} else if dia == 2 {
		fmt.Println("El dia es Martes.")
	} else if dia == 3 {
		fmt.Println("El dia es Miercoles.")
	} else if dia == 4 {
		fmt.Println("El dia es Jueves.")
	} else if dia == 5 {
		fmt.Println("El dia es Viernes")
	} else if dia == 6 {
		fmt.Println("El dia es Sabado.")
	} else if dia == 7 {
		fmt.Println("El dia es Domingo.")
	} else {
		fmt.Println("El dia que ingreso no existe.")
	}

	switch dia {
	case 1:
		fmt.Println("El dia es Lunes.")
	case 2:
		fmt.Println("El dia es Martes.")
	case 3:
		fmt.Println("El dia es Miercoles.")
	case 4:
		fmt.Println("El dia es Jueves.")
	case 5:
		fmt.Println("El dia es Viernes")
	case 6:
		fmt.Println("El dia es Sabado.")
	case 7:
		fmt.Println("El dia es Domingo.")
	default:
		fmt.Println("El dia que ingreso no existe.")
	}

	numero := 26
	switch {
	case numero > 24:
		fmt.Println("El", numero, "es mayor que 24")
		fallthrough
	case numero > 25:
		fmt.Println("El", numero, "numero es mayor que 25")
	}
}
