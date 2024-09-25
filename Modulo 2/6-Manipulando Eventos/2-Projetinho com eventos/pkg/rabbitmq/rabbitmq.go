package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Cria uma conecao e um canal(channeel) para enviar e receber mensagens do rabbitmq
func OpenChannel() (*amqp.Channel, error) {
	// cria uma conexao com o rabbitmq
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err) //mata a conexao se der erro
	}
	ch, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

// Publish/publicar envia uma mensagem para o rabbitmq
func Publish(ch *amqp.Channel, body string, exchangeName string) error {
	//OBS: temos que ligar o amq.direct em uma fila no rabbitmq para a mensagem ser enviada para a fila
	//É necessario cria uma ou mais filas no site do rabbitmq

	err := ch.Publish(
		exchangeName, // exchange - nome do exchange (se nao tiver, envia para a fila default) //estou passando o valor amq.direct é o nome do exchange
		"",           // routing key - nome da fila que quero enviar a mensagem
		false,        // mandatory - se a mensagem é obrigatoria (nao é usado com frequencia)
		false,        // immediate - se a mensagem deve ser enviada imediatamente (nao é usado com frequencia)
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Consume/Consumir recebe uma mensagem de uma fila do rabbitmq
func Consume(ch *amqp.Channel, out chan amqp.Delivery, queue string) error {
	msgs, err := ch.Consume( //Cria um consumer para consumir a fila
		queue,         // nome da fila que quero ler
		"go-consumer", // consumer - nome da aplicação que esta consumindo a fila
		false,         // auto-ack - se a mensagem foi consumida com sucesso, ela é removida da fila (so fazemos isso se tivermos certeza que a mensagem foi processada com sucesso e que nao vamos mais prdecisar)
		false,         // exclusive - se a fila é exclusiva para o consumer (se for true, a fila é exclusiva para o consumer que a criou e é excluida quando a conexao com o rabbitmq é fechada)
		false,         // no-local - se o rabbitmq nao deve enviar mensagens para o mesmo consumer que as publicou (nao é usado com frequencia)
		false,         // no-wait - se o rabbitmq deve esperar uma resposta do consumer (nao é usado com frequencia)
		nil,           // args - argumentos adicionais (nao é usado com frequencia)
	)
	if err != nil {
		return err
	}
	//Consome a fila
	for msg := range msgs {
		out <- msg //joga a mensagem no canal do Go
	}
	return nil
}
