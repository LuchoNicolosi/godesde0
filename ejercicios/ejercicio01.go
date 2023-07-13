package ejercicios

import (
	"fmt"
	"strconv"
)

func ConviertoAentero(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if num > 100 {
		fmt.Println("Es mayor a 100.")
	} else {
		fmt.Println("Es menor a 100.")
	}

	if err!=nil {
		fmt.Println(err)
	}
	
	return num, texto
}
