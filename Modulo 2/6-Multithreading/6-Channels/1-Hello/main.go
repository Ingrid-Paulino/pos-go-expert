package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Cria um canal de strings (Vazio)

	//Trhead 2
	go func() {
		canal <- "Olá Mundo!" // Envia a mensagem para o canal (Bloqueante, Está cheio)
	}()

	msg := <-canal // Recebe a mensagem do canal (Esvazia o canal)
	fmt.Println(msg)
}
