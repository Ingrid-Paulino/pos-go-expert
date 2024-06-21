package main

import (
	"fmt"
)

func main() {

	// Closures: são funções anonimas, ou seja, dentro de uma funcao, vc pode ter outra funcao
	func() int {
		return sum(1, 5, 7, 9) * 2
	}()

	//Posso colocar dentro de uma variavel
	total := func() int {
		return sum(10, 50, 70, 90) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
