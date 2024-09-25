package main

import (
	"fmt"
	"net/http"
)

// conta a quantidade de visitar que o meu site teve
var number uint64 = 0

// Conflito de Concorrencia:
// Esse servidor pra cada requisicão "curl localhost:3000/" o Go cria uma nova thread, isso pode gerar concorrencia na variavel number
// Podemos testar isso rodando varias solicitacoes com o apache Bench que é uma ferramenta para teste de carga e benchmarking (tem que instalar no pc, no mac ja tem por default)
// comando: ab -n 10000 -c 100 http://localhost:3000 (10000 é o numero de solicitacoes que vao rodar, 100 é a qtd de pessoas que vao acessar ao mesmo tempo)
// rode curl localhost:3000/ e vc verá que o numero de visitas é inferior do que solicitamos no apache bench por causa da concorrencia
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes", number)))
	})

	http.ListenAndServe(":3000", nil)
}

// O Go permite a gente descobrir se o nosso programa esta tendo problema de concorrencia
// COMANDO: go run -race main.go (verifica se temos algum senario de race condition no nosso programa)
// curl localhost:3000/ rode 1
// curl localhost:3000/ rode 2
// curl localhost:3000/ rode 3
// ab -n 50000 -c 100 http://localhost:3000/ - teste de carga (podemos perceber que onde rodamos go run -race main.go deu erro de WARNING: DATA RACE)

//OBS: solucao de como resolver esse problema na AULA 5
