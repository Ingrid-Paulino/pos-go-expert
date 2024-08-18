package main

import (
	"fmt"

	"github.com/ingrid-paulino/fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

//OBS: Paaso a passo :
//Rode o docker: docker-compose up -d
//Confirma a conexao no terminal ou no browser: http://localhost:15672/
//Rode o projeto com o comando go run main.go para abrir uma conexao no rabbitmq

//Teste mais rapido para publicar msg pelo proprio rabbitmq e aparecer no terminal o retorno :
/*
- rodando os comandos a cima, é posivel criar um topico na aba Queues no rabbitmq
- depois de criado clique em cima dele eé posivel publicar uma msg pela pagina web na aba publish message
- pare o programa no terminal e rode de novo. A msg publicada tem que ser consumida(aparecer no terminal)
*/

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "orders") //orders é o nome da fila que quero consumir la do rabbitmq
	for msg := range msgs {                 //pega todas as msgs que vai chegando no canal (rabbitmq.Consumer(ch, msgs)
		fmt.Println(string(msg.Body))
		msg.Ack(false) // coloco o mesmo valor que configurei o auto-ack no consumer
	}
}
