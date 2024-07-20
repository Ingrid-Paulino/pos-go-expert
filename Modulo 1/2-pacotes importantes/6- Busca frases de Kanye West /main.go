package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Frase struct {
	Quote string `json:"quote"`
}

func main() {
	//rode: go run main.go ola -> imprime ola no terminal
	// for _, url := range os.Args[1:] { //os.Args[1:] : Retorna um slice com letra a letra do que eu digitar e imprimi
	// 	println(url)
	// }

	//Roda o programa: go run main.go https://api.kanye.rest
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err) //Fprintf: joga o resultado pra algum lugar, nesse caso vai jogar para o os.Stderr
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data Frase
		err = json.Unmarshal(res, &data) //transforma em struct
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println(data)
	}
}
