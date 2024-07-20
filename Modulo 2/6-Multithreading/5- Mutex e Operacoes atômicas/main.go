package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// conta a quantidade de visitas que o site teve
var number uint64 = 0

// Resolvendo conflito de Concorrencia da Aula 4:
// Podemos testar isso rodando varias solicitacoes com o apache Bench que é uma ferramenta para teste de carga e benchmarking (tem que instalar no pc, no mac ja tem por default)
// comando: ab -n 10000 -c 100 http://localhost:3000 (10000 é o numero de solicitacoes que vao rodar, 100 é a qtd de pessoas que vao acessar ao mesmo tempo)
// rode curl localhost:3000/ e vc verá que o numero de visitas é 10001

// Forma 1 de resolver é com Mutex(lock e unlock)

func main() {
	var m sync.Mutex
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock() //quando alguem acessar para mudar o valor dessa variavel, será bloqueado para os outros acessos e ningem mais vai poder acessar essa variavel
		number++
		m.Unlock()                        //quando terminar de mudar o valor da variavel, será desbloqueado para os outros acessos e os outros vao poder acessar essa variavel
		time.Sleep(30 * time.Millisecond) //-- todo: com a time funcionou, sem nao funcionou - segue com problema de concorrencia
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes\n", number)))
	})

	http.ListenAndServe(":3000", nil)
}

// Forma 2 de resolver é com soma atomica
// usando o pacote atomic, pacote proprio do GO para resolver problemas de concorrencia
//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		//atomic faz o processo de lock e unlock de forma automatica
//		//atomic tem varias outras funcoes, todo ler a doc atomic
//		atomic.AddUint64(&number, 1) //adiciona 1 ao valor da variavel number (atomic.AddUint64 é uma operacao atomica, ou seja, é uma operacao que nao pode ser interrompida por outra thread, ela é executada de forma atômica, ou seja, ela é executada de uma vez só, sem interrupções. Isso garante que a operação seja concluída antes que outra thread possa acessar a variável.
//		//time.Sleep(300 * time.Millisecond) //todo com a time funcionou, sem n funcionou - segue com problema de concorrencia
//		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes", number)))
//	})
//
//	http.ListenAndServe(":3000", nil)
//}
