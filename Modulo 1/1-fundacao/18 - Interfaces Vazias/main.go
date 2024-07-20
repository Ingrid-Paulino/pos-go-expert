package main

import "fmt"

// Interface vazia implementa todo mundo
// Era muito usado quando não tinha generics
// Tentar evitar usar isso pois a pegada do GO é ter tipagem porte
type x interface{}

func main() {

	//Interface vazia aceita qualquer tipo
	var x1 interface{} = 10
	var y1 interface{} = "Hello, World!"

	showType(x1)
	showType(y1)

}

func showType(t interface{}) { //a função recebe qualquer tipo como parametro
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
