package main

import (
	"fmt"

	"github.com/luchonicolosi/godesde0/ejercicios"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoAtexto(1599)
	// fmt.Println(estado)
	// fmt.Println(texto)
	// if SO := runtime.GOOS; SO == "Linux." || SO == "OS X." {
	// 	fmt.Println("Esto no es windows, es", SO)
	// } else {
	// 	fmt.Println("Esto es windows :", SO)
	// }

	// switch SO := runtime.GOOS; SO {
	// case "Linux":
	// 	fmt.Println("Esto es linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", SO)
	// }

	num, texto := ejercicios.ConviertoAentero("sadasd")

	fmt.Printf("Texo: %s\nNumero: %d", texto, num)
}