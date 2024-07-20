package main

import (
	"fmt"
	"log"
	"net/http"
)

/*Criação de um midlewere para servidor web que toda vez que receber um panic, vamos recuperar e o servidor web e não vai cair
poderiamos fazer isso pra diversas coisas, ate mesmo pra uma conecção de banco ...*/

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	log.Println("Listening on", ":2000")
	if err := http.ListenAndServe(":2000", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", ":2000", err)
	}
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entrou no handler")
		//recuperamos o panic e nn deixamos o servidor cair
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered panic: %v\n", r)
				//debug.PrintStack() //Recupera panics, aplicação vai continuar rudando, porem vai logar o panic
				http.Error(w, "Internal server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
