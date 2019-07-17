package main

import "fmt"

func main() {
	var numero int
	numero = 25
	fmt.Println(numero)
	numero = 45
	fmt.Println(numero)

	var nombre string
	nombre = "Victor"
	nombre_dos := "Alejandro"
	fmt.Println(nombre, nombre_dos)

	nombre, nombre_dos = nombre_dos, nombre
	fmt.Println(nombre, nombre_dos)

	nombre_tres, nombre := "Maria", "Pedro"
	fmt.Println(nombre_tres, nombre)
}
