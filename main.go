package main

import (
	"fmt"

	"github.com/luchonicolosi/godesde0/variables"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoAtexto(1599)
	fmt.Println(estado)
	fmt.Println(texto)
}
