package arreglos_slices

import "fmt"

//var tabla []int --> Esto es un slice, no tiene una cantidad de dimension de elementos
var tabla [10]int = [10]int{10, 0, 59} //Esto es un vector, le hardcodeo la cantidad de dimension de elementos
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 32
	tabla[2] = 10

	tabla2 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
