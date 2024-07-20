package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() //iniciando o contexto (esta em branco)
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	//select: trabalha de forma assincrona, ele fica aguardando os resultados e quando o resultado chega ele toma a ação
	select {
	case <-ctx.Done(): //caso o contexto for done (<- === receve) -> case recebe conteco finalizado. O done acontece aqui quando passar o meu timeout de 3 segundo
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(2 * time.Second): //caso passe 5 segundos e não foi cancelado o contexto vai implimir a msg. Esse time.after poderia ser qualquer coisa, ate mesmo uma chamada a api
		fmt.Println("Hotel booked.")
	}
}

//TODO: escrever certinho depois
/*Funções importantes do context:
WithCancel: pode ser cancelado em qualquer momento independente de tempo
WithDeadline: tem um tempo de daqui quanto tempo pode ser cancelado
WithTimeout: é como se fosse uma contagem regresiva
WithValue: guarda valor*/
