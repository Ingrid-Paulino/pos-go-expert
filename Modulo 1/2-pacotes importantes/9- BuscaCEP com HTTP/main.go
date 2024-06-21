package main

import (
	"encoding/json"
	"io"
	"net/http"
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

// Para rodar use a extençao Thunder client do vscode: GET http://localhost:8080?cep=31330500 - curl http://localhost:8080?cep=31330500
func main() {

	http.HandleFunc("/", BuscaCEPHandler)

	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	//Headers: tem os seguintes metodos add, get, set, get values, has (confirmar se existe valores), Write
	//ResponseWriter: tem os seguintes metodos Header, Write (quando respondemos alguma coisa no status http), WriteHeader  (respondemos com statusCode)
	if r.URL.Path != "/" { //tratativa
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep") //manipulação de query string -> pega valor por parametro na url. Rode: (http://localhost:8080/?cep=31330500)
	if cepParam == "" {                  //tratativa
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := BuscaCEP(cepParam)
	if err != nil { //tratativa
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") //indica que vai retornar no formato json
	w.WriteHeader(http.StatusOK)                       // retorna status code

	//Forma 1 - usado mais quando queremos o resposta em uma variavel para fazer outras manipulaçoes
	// result, err := json.Marshal(cep)
	// if err != nil { //tratativa
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)
	//Forma 2 - resposta direta
	json.NewEncoder(w).Encode(cep) //converte struct pra json e joga no http.ResponseWriter
}

func BuscaCEP(cep string) (*ViaCEP, error) { //nn precisava retornar um ponteiro de ViaCEP, o wesley retornou por custume e por trabalhar no dia a dia com o valor real e não com copias
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()            //fecha a conecção do servidor para não ter vazamento de recursos
	body, err := io.ReadAll(resp.Body) //recebo valor em json
	if err != nil {
		return nil, err
	}

	var c ViaCEP
	err = json.Unmarshal(body, &c) //trensformo valor json em struct
	if err != nil {
		return nil, err
	}

	return &c, nil
}
