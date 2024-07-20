package main

import "fmt"

//é uma boa pratica setar na tipagem a direção do canal
// chan <- string: canal recebe informaçoes - receive only
// <-chan string: canal envia/entrega informaçoes - send only

func recebe(nome string, hello chan<- string) { //chan <- string: sig que estou setando a direção do canal, nesse caso só posso enviar dados para o canal - sig que o canal só pode receber informaçoes
	hello <- nome // Envia a mensagem para o canal (Bloqueante, Está cheio)
}

func ler(data <-chan string) { //<-chan string: sig que estou setando a direção do canal, nesse caso o canal só ira entregar resultados
	fmt.Println(<-data) // Recebe a mensagem do canal (Esvazia o canal)
}

// Thread 1
func main() {
	// o canal sempre vai receber um dado ou enviar um dado
	hello := make(chan string) // Cria um canal de strings (Vazio)
	go recebe("Hello", hello)
	ler(hello)
}
