package main

import (
	"fmt"
	"time"
)

// A ideia do select é ser um switch case para canais
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c1 <- 1
		time.Sleep(time.Second * 4)
	}()
	go func() {
		time.Sleep(time.Second * 5)
		c2 <- 2
	}()

	//For vai rodar duas vezes e retornar dois valores
	for i := 0; i < 2; i++ {
		//Com Select conseguimos imprimir o valor do canal que chegar primeiro
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 3): //se demorar mais de 3 segundos, o select vai executar esse case
			fmt.Println("timeout")
		default:
			fmt.Println("default") //Vai rodar primeiro que todos os outros cases
		}
	}

}
