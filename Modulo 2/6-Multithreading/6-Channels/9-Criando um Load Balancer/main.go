package main

import (
	"fmt"
	"time"
)

//Load Balancer: é um padrão de projeto que distribui o trabalho entre vários processadores, permitindo que um processo seja executado em paralelo em vários processadores.
//Load Balancer: é um balanceador de carga. Processa grande quantidades de dados de forma paralela e distribuida, para que o processamento seja mais rapido e eficiente.
/*
Ex de um servidor web: Vamos supor que vamos receber um monte de requisicoes de usuarios, para receber em parelelo e distrbuido de forma simultanea sem ter que travar o meu processo
podemos usar um load balancer para distribuir as requisições entre vários servidores, para que o processamento seja mais rápido e eficiente.
p
*/

// vai processar os dados
func worker(workwerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d: Received %d\n", workwerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	//forma 1
	//Forma serial: um unico worker
	//go worker(1, data) // faz de forma serial, vai rodar um por um. As chamadas são enfileiradas

	//forma 2
	//Forma paralela: varios workers - load balancer
	//podemos evitar esse tipo de problema e ter um processamento mais rapido usando load balancer, criando varios workers
	//go worker(1, data)
	//go worker(2, data)
	//go worker(3, data)

	//forma 3 -> usando um loop para criar varios workers de forma dinamica
	qtdWorkers := 100 //quantidade de workers que vamos criar
	//inicializa os workers
	for i := 0; i < qtdWorkers; i++ { // dessa forma meu programa vai rodar mais rapido, pois vai rodar de forma paralela
		go worker(i, data)
	}

	for i := 0; i < 1000; i++ {
		data <- i // envia informaçoes para o canal
	}
}
