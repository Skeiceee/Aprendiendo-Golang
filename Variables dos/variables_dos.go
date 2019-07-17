package main

import "fmt"

func main() {
	//Nombrar variables
	numero := 25
	_nombre := "Alejandro"
	nombreUsuario := "Qwerty"
	fmt.Println(numero, _nombre, nombreUsuario)

	var (
		dios               = "Goku"
		enemigo1, enemigo2 = "Babidi", "Cell"
	)
	fmt.Println(dios, enemigo1, enemigo2)

	var numero_dos int
	numero_dos = 30
	fmt.Println(numero_dos)

	//Scope
	var razaGoku = "Saiyajin"
	imprimir(razaGoku)
}

func imprimir(razaGoku string) {
	fmt.Println("La raza de goku es: " + razaGoku)
}
