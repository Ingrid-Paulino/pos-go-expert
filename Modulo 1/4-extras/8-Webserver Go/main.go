package main

//Novidades e correçoes de bud da versão 1.22 do go

import (
	"fmt"
	"net/http"
)

func main() {
	//melhoria do servidor web, os parametros pasados por {} na url, podem ser recuperados nos metodos

	mux := http.NewServeMux()
	//mux.HandleFunc("GET /books/{id}", GetBookHandler)

	//{d...}: pega varias chamadas independente da quantidade de recursos - GET http://localhost:9000/books/dir/djdjdjd/aa/add/dddre/gfgf
	// mux.HandleFunc("GET /books/dir/{d...}", BooksPathHandler)

	// pega uma rota exata //{$}: indica que qualquer coisa pra frente de /books não sera pego
	// mux.HandleFunc("GET /books/{$}", BooksHandler)

	//essa rota é mais especifica, entt tem precedencia da debaixo que finaliza com {x}.
	// mux.HandleFunc("GET /books/precedence/latest", BooksPrecedenceHandler)
	// mux.HandleFunc("GET /books/precedence/{x}", BooksPrecedence2Handler)

	//Ambas as rotas dão match no path, por isso o go solta um erro pois o nivel de expecificidade de ambos é o mesmo
	//mux.HandleFunc("GET /books/{s}", BooksPrecedenceHandler)
	//mux.HandleFunc("GET /{s}/latest", BooksPrecedence2Handler)
	http.ListenAndServe(":9000", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book " + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirpath := r.PathValue("d") // Access captured directory path segments as slice
	fmt.Fprintf(w, "Accessing directory path: %s\n", dirpath)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}

func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence"))
}

func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence 2"))
}
