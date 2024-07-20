package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	//rode: go run main.go ola -> imprime ola no terminal
	// for _, url := range os.Args[1:] { //os.Args[1:] : Retorna um slice com letra a letra do que eu digitar e imprimi
	// 	println(url)
	// }

	//Comando para rodar: go run main.go http://viacep.com.br/ws/31330-500/json/
	// for _, url := range os.Args[1:] {
	for _, cep := range os.Args[1:] {
		// req, err := http.Get(url)
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/") // o run main.go 31330-500 -> poderia passar a url de uma vez e mandar somente o cep no terminal

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err) //Fprintf: joga o resultado pra algum lugar, nesse caso vai jogar para o os.Stderr
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data) //transforma em struct
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println(data) //retorno: {31330-500 Rua José Ribeiro Filho  Ouro Preto Belo Horizonte MG 3106200  31 4123}
		fmt.Println(data.Localidade)

		//Cria arquivo
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		//grava dados de retorno no arquivo
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf)) //fmt.Sprintf: permite que pegamos a saida e jogue em algum lugar
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gravar no arquivo: %v\n", err)
		}
		fmt.Println("Arquivo criado com sucesso")
		fmt.Println("Cidade: ", data.Localidade)

		//caso queira ver o resultado pelo terminal, rode:
		//ls
		//cat nomeDoArquivo/cidade.txt

		//buildar:
		// go build -o cep main.go -> criou o arquivo cep
		//para rodar: ./cep 31330500
	}
}
