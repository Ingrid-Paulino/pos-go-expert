package main

import "net/http"

//Para rodar:
//1: subir o servidor em um terminal: "go run main.go"
//2: em outro terminhar chamar a rota: curl localhost:8080/?cep=31330500
//o passo 2 no terminal aparece so o texto as tratativas com retorno apenas com statusCode não aparece
//3: para ver melhor podemos rodar no postman, insominia ou no browser: (localhost:8080(status 400) - http://localhost:8080/?cep=31330500(status 200), http://localhost:8080/hello(status 404))

func main() {

	http.HandleFunc("/", BuscaCEPHandler)

	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { //tratativa
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep") //manipulação de query string -> pega valor por parametro na url. Rode: (http://localhost:8080/?cep=31330500)
	if cepParam == "" {                  //tratativa
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") //indica que vai retornar no formato json
	w.WriteHeader(http.StatusOK)                       // retorna status code
	w.Write([]byte("Hello, World!\n"))                 //retorna uma mensagem de slice de bytes
}
