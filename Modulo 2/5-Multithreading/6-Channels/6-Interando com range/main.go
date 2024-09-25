package main

import "fmt"

// sincornização de dados entre duas threads diferentes
// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch) //Se essa tread for uma goroutine, ela vai terminar antes de receber a mensagem do canal e meu programa vai morrer
}

func reader(ch chan int) {
	for x := range ch { // Recebe a mensagem do canal e esvaia o canal
		fmt.Printf("Received %d\n", x)
	}
	//OBS: O range só sai do loop quando o canal estiver fechado na func publish, ao contrario dara deadlock
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ { //Toda vez que o for roda, ele envia o valor para o canal
		ch <- i //em quanto o canal estiver cheio, ele vai bloquear a execução e nao vai ser posivel enviar outro valor para o canal
	}
	close(ch) //Fecha o canal para que o range da função reader saia do loop e nao dar deadlock
}
