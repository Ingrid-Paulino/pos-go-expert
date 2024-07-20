package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

/*
Pacote de Contexto do GO : é um pacote que permite que passamos as informaçoes dele para diversas chamadas no nosso sistema.
Temos a opção que esses contextos sejam cancelados, e quando um contexto é cancelado, a operação para na hora,
pra que não fique gastando muito tempo
*/
func main() {
	ctx := context.Background() //contexto vazio
	//OBS: Usando o contexto nn precisamos usar o sistema de http.Client{Timeout: time.Second} -> esse timeout no final das contas estava setando o tempo do contexto
	//Nesse caso eu que estou criando o meu contexto
	ctx, cancel := context.WithTimeout(ctx, time.Second) // qualquer coisa que estiver usando esse contexto e passar de 5 seg, não vai continuar tentando, e vai parar/cancelar a execução.
	//ctx, cancel := context.WithTimeout(ctx, time.Millisecond) //da timeout
	//cancel() //segunda forma de cancelar um contexto
	defer cancel() // no final o contexto vai ser cancelado, ou vai ser pelo tempo da linha 13 ou ao final da execução do sistema

	//Forma 3
	//ctx, cancel := context.WithCancel(ctx) //essa func obrigatoriamente exige que seja rodado a func cancel() para cancelar um contexto

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil) //na configuração do contexto temos 1 seg pra ter a resposta dessa req
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
