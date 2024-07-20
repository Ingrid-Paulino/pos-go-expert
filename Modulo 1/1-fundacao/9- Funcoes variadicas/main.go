package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 5))
}

// funcoes variadicas: é usada quando nao sabemos a quantidade de parametros que vamos receber
// vamos somar uma variedade de numeros, mas n sei a quantidade de numeros que vou receber por paramentro
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
