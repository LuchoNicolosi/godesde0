package mapas

import "fmt"

func MostrarMapas() {
	//Como crear un map
	paises := make(map[string]string)

	futbolArgentino := map[string]int{
		"River":  70,
		"Racing": 29,
	}
	//Como asignarle key:value
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	//Como recorrer un map
	for equipo, titulos := range futbolArgentino {
		fmt.Printf("%s, tiene %d titulos.\n", equipo, titulos)
	}

	//Como borrarle elementos
	delete(futbolArgentino, "Racing")

	//Actualizarlos
	futbolArgentino["River"] = 69

	fmt.Println(futbolArgentino)

	// Como preguntar si una clave existe, y obtener su valor
	puntaje, existe := futbolArgentino["River"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe = %t \n", puntaje, existe)
}
