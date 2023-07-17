package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {
	var numero int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un numero : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	// fmt.Printf("------ La Tabla Del [%d] ------\n", numero)
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("[%dx%d] = %d\n", numero, i, numero*i)
	}

	return texto
}
