package main

import "fmt"

func main() {
	//Memória(go cria) -> Endereço(a) -> Valor(10)
	// "a" tem um endereço de memoria que o go cria
	// variavel -> aponta pra um ponteiro que tem um endereço na memoria -> que tem um valor
	a := 10
	fmt.Println(a)  // mostra o valor de 10
	fmt.Println(&a) //mostra o endereço da memoria de onde a está sendo guardado

	//Asterisco indica que vamos receber um endereço de memoria
	var ponteiro *int = &a
	fmt.Println(ponteiro) //retorna o endereço da memoria

	*ponteiro = 20 //mudo valor da memoria de a
	fmt.Println(a)

	b := &a         //b vai apontar pro mesmo endereço de memoria de a
	fmt.Println(b)  //printa o valor da memoria de a
	fmt.Println(*b) //printa o valor que esta guardado na memoria de a

	*b = 30 //muda o valor de a
	fmt.Println(a)
	fmt.Println(&b) //valor da memoria de b
}
