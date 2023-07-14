package ejercicios

import (
	"strconv"
)

func ConviertoAentero(texto string) (int, string) {
	num, _ := strconv.Atoi(texto)
	if num > 100 {
		return num, "Es mayor a 100."
	} else {
		return num, "Es menor a 100."
	}
}
