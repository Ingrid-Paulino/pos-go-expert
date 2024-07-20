package main

import (
	"log"
	"net/http"
	"time"
)

// para rodar: go run main.go e curl localhost:8080
/*se eu der um ctrl + C nesse comando curl localhost:8080, vai imprmir Request finalizada*/
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() //A request tem um context
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second): //poderia ser outro tipo de operação, ex: chamada a uma api ...
		//Imprime no comand line stdout
		log.Println("Request processada com sucesso") //vai imprimir na tela do servidor
		// Imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done(): //contexto cancelado/finalizado
		//Imprime no comand line stdout
		log.Println("Request cancelada pelo cliente") //vai imprimir na tela do servidor
	}
}

/*Se eu n der um case <- ctx.Done e o cliente cancelar, o programa vai continuar executando o processamento
por exemplo da linha 21, simplesmente pq não tenho um contexto. Isso evita fazer processamentos desnecessarios
conforme o tempo passa. Poderiamos estar processando n coisas ao mesmo tempo(ex: banco de dados...), o usuario deu ctrl + C
com o case <-ctx.Done() vai parar tudo pois cancelou o contexto e não tem o pq continuar. */
