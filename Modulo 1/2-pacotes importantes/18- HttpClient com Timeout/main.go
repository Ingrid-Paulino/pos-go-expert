package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*Quando estamos trabalhando em um sistema e ele tem que funcionar
de uma forma muito rapida, a primeira coisa que temos que pensar
é como deixar mais performatico posivel

- uma das formas é conseguindo estabelecer limites, limites de chamadas externas que o sistema vai realizar

ex: tenho um sistema de GO mais rapido do mundo, mas ele vai chamar uma api externa e essa api demora 10 seg pra retornar o resultado.
Isso vai fazer diferença na performace da minha aplicação.

- Uma forma de tratar isso, é colocando limites, caso a api nao responda em X tempo retorna um erro.
*/

func main() {
	c := http.Client{Timeout: time.Second} //Timeout é o tempo de duração maxima que a aplicação pode demorar //tempo limite que a aplicacao vai fazer a requisição
	//c := http.Client{Timeout: time.Microsecond} //estora o temo limite (timeout) e retorna um erro
	resp, err := c.Get("http://google.com")
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
