package variables

import (
	"fmt"
	"time"
	"strconv"
)

var nombre string //esta variable estara disponible para todos los archivos del paquete (variables)
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	//MuestroEnteros()
	nombre = "Lucho"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println(nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoAtexto(numero int) (bool, string) {
texto := strconv.Itoa(numero)
return true, texto
}
