package main

import (
	"fmt"
	"sync"
)

// OBS: Nesse caso não seria necessario usar o waitGroup, pois o range já faz o controle de quantas vezes o loop vai rodar, mas exixte muito usar waitGroup com canais
// Se não usassemos o waitgroup, teriamos que tirar a goroutine do reader e colocar o close do canal no final da função publish (aula anterior explica)
// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10) //Vamos interar por 10 vezes e vai rodar 10 vezes o wg.Done()

	//As duas funções vão rodar em paralelo em background e mesmo assim o programa vai esperar as duas terminarem para finalizar a execução por causa do waitGroup
	go publish(ch)
	go reader(ch, &wg)
	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch { // Recebe a mensagem do canal e esvaia o canal
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
	//OBS: O range só sai do loop quando o canal estiver fechado na func publish, ao contrario dara deadlock
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ { //Toda vez que o for roda, ele envia o valor para o canal
		ch <- i //em quanto o canal estiver cheio, ele vai bloquear a execução e nao vai ser posivel enviar outro valor para o canal
	}
	close(ch) //Fecha o canal para que o range da função reader saia do loop e nao dar deadlock
}
