package main

import "github.com/ingrid-paulino/fcutils/pkg/rabbitmq"

//OBS: Paaso a passo :
//Rode o docker: docker-compose up -d
//Confirma a conexao no terminal ou no browser: http://localhost:15672/
//Rode o projeto com o comando go run main.go para abrir uma conexao no rabbitmq
// Na aba Queues do rabbitmq, é posivel ver a fila que foi criada

/*
Para enviar uma msg daqui para uma fila é necessario criar uma fila no site do rabbitmq:
Passo a passo:
- faça a conexao com o rabbitmq e o login
- Na aba Exchange, clique em aqm.direct e crie uma fila na abinha Bindings
  - se eu n adicionar uma key na fila, a msg vai para a fila de forma default
*/

//OBS: deixe o consumer rodando e depois rode o producer para enviar a msg

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World!", "amq.direct") //amq.direct é o nome do exchange
}
