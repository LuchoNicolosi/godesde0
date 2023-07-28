package arreglos_slices

import "fmt"

var tablaSlice []int = []int{0, 1, 2, 3, 4, 5}
var arreglo [10]int = [10]int{1, 2, 3, 4, 21, 34, 5}

func MuestroSlice() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:]   //Slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  //Slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] //Slice creado desde un vector, desde la posicion 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)                                      //No solo sirve para slices!!, para todo tipo de estructura
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos)) //Largo 5, Capacidad 20

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums)) //Largo 100, Capacidad 128
}
