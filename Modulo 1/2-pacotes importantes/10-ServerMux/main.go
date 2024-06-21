package main

import (
	"net/http"
)

// para rodar: 1: go rum main.go \\ 2: curl localhost:8080/
func main() {

	//FORMA 1
	// http.HandleFunc("/", BuscaCEPHandler)
	//multiplax padrao
	// http.ListenAndServe(":8080", nil) //forma padrão, n temos q se preocupar em criar um http.NewServeMux (Por um lado é bom, mas pelo outro vc n tem controlle do mux, sendo assim, diversos lugares do sistema pode injetar endpoints que vc n queira. Outro ponto, se quiser criar diversos servidores, nn será possivel, pois o mux padrao toma conta de tudo)

	//FORMA 2
	//responsavel por atachar os hendlers
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// }) //estou adcionando esse handle no meu mux

	// http.ListenAndServe(":8080", mux)

	//Forma 3
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", HomeHandler)
	// mux.Handle("/blog", blog{}) //quando o meu servidor chamar /blog, ele vai executar o serveHTTP

	// http.ListenAndServe(":8080", mux)

	//Forma 4
	mux := http.NewServeMux() //com serverMux temos mais controle das nossas rotas
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My blog"}) //Dessa forma tem a flexibilidade de jogar valores para o metodo

	http.ListenAndServe(":8080", mux)

	//posso ter varios servidores
	mux2 := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Ingrid"))
	})
	http.ListenAndServe(":8081", mux2) //TODO mandar no forum pq nn funconou essa rota

	//A recomendação é sempre criar seu proprio mux e atachar o que quiser
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

//Forma 3
// type blog struct {}

// func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
// 	w.Write([]byte("Blog"))
// }

// Forma 4
type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
