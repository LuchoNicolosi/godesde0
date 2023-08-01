package ejer_interfaces

import (
	"fmt"

	"github.com/luchonicolosi/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respondando \n", hu.Sexo())
}

func SeresVivos(ser interfaces.SerVivo) {

	if ser.EstaVivo() {
		fmt.Println("Si, soy un ser vivo")
	}
}
