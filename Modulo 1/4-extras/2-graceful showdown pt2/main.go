package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*TEMA: Forma da aplicação parar de aceitar novas requisiçoes e começar a desligar o sistema caso aconteça
algum problema, para esperar as requisiçoes sendo processadas ainda acontecer*/

// run: go run main.go (parar a aplicação com crlt + c para ver a tratativa) e curl http://localhost:2000
func main() {
	server := &http.Server{Addr: ":2000"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*Parte 1 aula anterior.

		Para converter isso podemos implementar o graceful shutdown. Toda vez que o sistema for derrubado, normalmente
		derrubamos com alguma introdução ou com alguns avisos no sistema operacional, sinalisando que o sistema operacional será desligado
		ou que esta sendo interrompido o programa. */
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World\n"))
	})

	go func() {
		fmt.Println("Server is running at http://localhost:3000")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1) /*canal do tipo os.Signal criado: sinal do sistema operacional*/

	/*Notifica o canal para esvaziar quando receber o sinal de notficação do sistema operacional*/
	/*Notfica o canal quando tiver alguma interrupção, ou quando tiver alguma
	syscall.SIGTERM do sistema operacional, ou a interrupção enviada pelo sistema operacional syscall.SIGINT*/
	/*Toda vez que eu notificar esse canal, o canal será esvaziado e o programa ira acabar.*/
	//Dessa forma garanto, esperar 5 seg para o servidor terminar de processar tudo
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop /*bloqueia o programa, para que essa thread fica parada nesse ponto*/

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //a dedline sera de 5 seg depois co contexo, ou será cancelado
	//ctx é cancelado depois de 5 seg
	defer cancel()

	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil { //quando o contexto for cancelado, server.Shutdown será executado
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err) //caso nn consiga desligar o servidor com o Shutdown
	}
	fmt.Println("Server stopped")
}

//TODO: brincar com isso fazendo essa tratativa para chamada de uma api
