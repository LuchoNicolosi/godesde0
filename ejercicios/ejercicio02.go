package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func IngresoNumero() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un numero : ")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			for {
				fmt.Println("Ingrese un numero valido : ")
				if scanner.Scan() {
					numero, err = strconv.Atoi(scanner.Text())
					if err == nil {
						break
					}
				}
			}
		}
	}

	fmt.Printf("------ La Tabla Del [%d] ------\n", numero)
	for i := 1; i <= 10; i++ {
		fmt.Printf("[%dx%d] = %d\n", numero, i, numero*i)
	}

}
