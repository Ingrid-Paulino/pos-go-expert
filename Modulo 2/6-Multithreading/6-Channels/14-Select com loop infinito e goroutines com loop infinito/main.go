package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

// A ideia do select é ser um switch case para canais
func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	//vamos supor que nessa thread o rabbitmq vai esta sendo processado
	go func() {
		for { //vai rodar pra sempre
			atomic.AddInt64(&i, 1) //incrementa o valor de i de forma segura para ser usado em goroutines concorrentes e nao ter problema de duas threads incrementarem o valor de i ao mesmo tempo
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()
	//vamos supor que nessa thread o kafka vai esta sendo processado
	go func() {
		for { //vai rodar pra sempre
			atomic.AddInt64(&i, 1) //incrementa o valor de i de forma segura para ser usado em goroutines concorrentes e nao ter problema de duas threads incrementarem o valor de i ao mesmo tempo
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	//Colocando o Select em um loop infinito para sempre ficar escutando os canais e executar o case que for chegando
	for { //vai rodar pra sempre
		//Com Select conseguimos imprimir o valor do canal que chegar primeiro
		select {
		case msg := <-c1: //rabbitmq
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg.id, msg.Msg)
		case msg := <-c2: //kafka
			fmt.Printf("Received from kafka: ID: %d - %s\n", msg.id, msg.Msg)
		case <-time.After(time.Second * 2): //se demorar mais de 3 segundos, o select vai executar esse case
			fmt.Println("timeout") //Vai ficar dando timeout a cada 3 segundos infinitamente
		}
	}
}
