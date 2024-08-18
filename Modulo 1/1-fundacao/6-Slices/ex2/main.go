package main

import "fmt"

func main() {
	evento := []string{"teste", "teste2", "teste3", "teste4"}
	fmt.Println(evento) //Imprime todos os valores: [teste teste2 teste3 teste4]
	//evento = evento[:2]
	//fmt.Println(evento) //Imprime os dois primeiros valores: [teste teste2]
	//evento = evento[2:]
	//fmt.Println(evento) //Imprime os dois ultimos valores: [teste3 teste4]
	//evento = evento[2:]
	//fmt.Println(evento) //Imprime um slice vazio: []
	//evento = evento[1:] //Remove o primeiro elemento do slice e mantem o restante
	//fmt.Println(evento)

	//Deleta o primeiro elemento do slice
	event := append(evento[:0], evento[1:]...) //O operador ... é usado para desempacotar o slice
	fmt.Println(event)

}
