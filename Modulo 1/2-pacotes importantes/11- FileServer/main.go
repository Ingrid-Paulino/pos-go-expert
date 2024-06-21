package main

import (
	"log"
	"net/http"
)

// para rodar: 1: go rum main.go \\ 2: no browser -> localhost:8080/ \\3: no browser -> localhost:8080/blog
func main() {
	//como poderia criar um servidor de arquivos para servir imagens, css, arquivo html...
	//podemos criar um fileServer e atachar no go

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from blog"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux)) //Se der algum erro vai ser logado
}
