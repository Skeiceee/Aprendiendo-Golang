package main

import (
	"fmt"
)

func main() {

	var (
		numero, numerosPar, numerosImpar = 0, 0, 0
	)

	fmt.Println("-----------------------------------------")
	fmt.Println("Discriminador de numeros pares e impares.")
	fmt.Println("-----------------------------------------")

	for i := 1; i <= 5; i++ {
		fmt.Print("Ingrese un numero: ")
		fmt.Scanln(&numero)
		if numero%2 == 0 {
			fmt.Println("El numero", numero, "es par")
			numerosPar++
		} else {
			fmt.Println("El numero", numero, "es impar")
			numerosImpar++
		}
	}

	fmt.Println("-----------------------------------------")
	fmt.Println("Hubieron", numerosPar, "numeros par.")
	fmt.Println("Hubieron", numerosImpar, "numeros impar.")
	fmt.Println("-----------------------------------------")

}
