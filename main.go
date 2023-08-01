package main

import (
	e "github.com/luchonicolosi/godesde0/ejer_interfaces"
	"github.com/luchonicolosi/godesde0/modelos"
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

	// num, texto := ejercicios.ConviertoAentero("50")

	// fmt.Printf("%s\nNumero: %d", texto, num)

	// teclado.IngresoNumero()

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.TablaMultiplicar())

	// files.GrabaTabla()
	// files.SumaTabla()
	// files.LeerArchivo()

	// funciones.Calculos()
	// funciones.LlamarClosure()
	//funciones.Exponencia(2)

	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlice()
	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()

	Pedro := new(modelos.Hombre)

	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer) //Soy un Hombre, y estoy respondando
	e.HumanosRespirando(Maria)  //Soy un Mujer, y estoy respondando

	e.SeresVivos(Maria)
	e.SeresVivos(Pedro)
}
