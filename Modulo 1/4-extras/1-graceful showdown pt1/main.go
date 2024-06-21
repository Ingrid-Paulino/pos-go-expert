package main

import (
	"net/http"
	"time"
)

/*TEMA: Forma da aplicação parar de aceitar novas requisiçoes e começar a desligar o sistema caso aconteça
algum problema, para esperar as requisiçoes sendo processadas ainda acontecer*/

// run: go run main.go e curl http://localhost:2000
func main() {
	server := &http.Server{Addr: ":2000"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*OBS: se eu parar o meu servidor com ctrl + c vou receber a
		mensagem "^Csignal: interrupt e curl: (52) Empty reply from server"
		Ele recebeu uma msg em branco, na hora que estava processando a requisição o servidor parou na metade ou poderia ter reiniciado
		e ferrou a requisiçao pois ela parou na metade.

		Para converter isso podemos implementar o graceful shutdown. AULA: 2-graceful showdown pt2 */
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World\n"))
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
