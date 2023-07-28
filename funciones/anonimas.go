package funciones

import "fmt"

func Calculos() {
	calculo := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println(calculo(10, 25))

	calculo = func(n1 int, n2 int) int {
		return n1 * n2
	}
	fmt.Println(calculo(10, 25))
}
