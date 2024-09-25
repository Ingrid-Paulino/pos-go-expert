package main

import (
	"fmt"
	"time"
)

type Message struct {
	id  int
	Msg string
}

// A ideia do select é ser um switch case para canais
func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	//vamos supor que nessa thread o rabbitmq vai esta sendo processado
	go func() {
		time.Sleep(time.Second * 2)
		msg := Message{1, "Hello from RabbitMQ"}
		c1 <- msg
	}()
	//vamos supor que nessa thread o kafka vai esta sendo processado
	go func() {
		time.Sleep(time.Second)
		msg := Message{1, "Hello from RabbitMQ"}
		c2 <- msg
	}()

	//Colocando o Select em um loop infinito para sempre ficar escutando os canais e executar o case que for chegando
	for {
		//Com Select conseguimos imprimir o valor do canal que chegar primeiro
		select {
		case msg := <-c1: //rabbitmq
			fmt.Printf("Received from RabbitMQ: %s\n", msg.Msg)
		case msg := <-c2: //kafka
			fmt.Printf("Received from kafka: %s\n", msg.Msg)
		case <-time.After(time.Second * 3): //se demorar mais de 3 segundos, o select vai executar esse case
			fmt.Println("timeout") //Vai ficar dando timeout a cada 3 segundos infinitamente
		}
	}
}
