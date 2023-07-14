package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 1; i <= 10; i++ {
		if i==6 {
			continue //continua, aunque i==6
			// break //corta la ejecucion
		}
		fmt.Println(i)
	}
}
