package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	// El defer Nos permite configurar cual va ser la ultima instruccion en ejecutarse cuando salga de la funcion
	fmt.Println("Este es un primer mensaje")

	defer fmt.Println("Este defer se ejecuta al final")
	defer fmt.Println("Este defer se ejecuta primero")

	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {
	//panic: Te permite abortar el sistema con un mensaje
	//recover: Te permite recuperarte de un panic (no hay forma de usar el recover si el defer)
	defer func() {
		reco := recover()
		if reco != nil {
		}
		log.Fatalf("Ocurrio un error, que genero Panic \n %v", reco)

	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
