package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Temos o Client que faz a requisição e temos a requisição. Ambos são coisas diferentes. Temos os dados da requisição e temos o client que usa essa requisição pra fazer chamadas
//Podemos fazer a separação e configurar a requisição da forma que quiser

func main() {
	//objeto de request e client http são coisas diferentes
	c := http.Client{}                                           //client http
	req, err := http.NewRequest("GET", "http://google.com", nil) //objeto de request
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	//junta o objeto da request + http client
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
