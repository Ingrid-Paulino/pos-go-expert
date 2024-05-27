package main

import "net/http"

func main() {

	//com função anonima - forma 1
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { //r tem todos os dados que recebo na requisição e w terá toda a resposta que quero entregar na request
		w.Write([]byte("Hello, World!\n"))
	})

	//forma 2
	http.HandleFunc("/", BuscaCEP)

	//para rodar:
	//1- go run main.go
	//2- curl localhost:8080
	//sobe o servidor http (somente esse comando teremos o retorno -> 404 page not found)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) { //r tem todos os dados que recebo na requisição e w terá toda a resposta que quero entregar na request
	w.Write([]byte("Hello, World!\n"))
}
