package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cadena string

	cadena = "Hola mundo!"

	fmt.Println(cadena)
	fmt.Println(len(cadena))
	fmt.Println(cadena[2])
	fmt.Println(cadena[:4]) // [0:4]

	cadena = cadena[:len(cadena)-1] + " nuevamente!"
	fmt.Println(cadena)

	cadena += "!!!!!"
	fmt.Println(cadena)

	cadena = `
		<html>
			<head>
				<meta charset="utf-8">
				<title>Welcome</title>
			</head>
			<body>
			</body>
		</html>
	`
	fmt.Println(cadena)

	cadena = "Hola Mundo \"25\""
	fmt.Println(cadena)

	edad := 21
	cadena = "Mi edad es " + strconv.Itoa(edad)
	fmt.Println(cadena)
}
