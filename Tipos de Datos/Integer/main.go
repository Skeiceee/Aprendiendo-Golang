package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//	Enteros con signos
	var entero8 int8
	// var entero16 int16
	var entero32 int32
	var entero64 int64

	//	Enteros sin signos
	// var uentero8 uint8
	// var uentero16 uint16
	// var uentero32 uint32
	// var uentero64 uint64

	//	Alias
	// var enteroByte byte
	var enteroRune rune

	//	Tipos dependientes de la plataforma
	// var enteroUint uint
	var enteroInt int
	// var enteroUintptr uintptr

	//	Conversión entre tipos

	//	Suma int32 y int64
	entero32 = 25 // int32
	entero64 = 30 // int64
	fmt.Println(entero32 + int32(entero64))

	//	Suma int32 y rune
	entero32 = 30
	enteroRune = 46
	fmt.Println(entero32 + enteroRune)

	//	Suma int32 y int
	entero32 = 50
	enteroInt = 100
	fmt.Println(entero32 + int32(enteroInt))

	fmt.Println(unsafe.Sizeof(enteroInt), unsafe.Sizeof(entero8))
}
