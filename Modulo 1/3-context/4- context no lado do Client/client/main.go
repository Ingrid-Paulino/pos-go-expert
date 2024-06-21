package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

/*OBS: esse exemplo esta ligado com a aula 3- context utilizando server http
É um exemplo para mostrar que podemos controlar o contexto do lado do cliente quanto do lado do servido
Sendo assim, conseguimos fazer o servidor parar de processar oque ele estava, já que o client se desconectou a ele (desconecta pois da timeout do context e cancelamos a operação)

Para rodar esse arquivo tem que rodar primeito o da aula 3- context utilizando server http e podemos brincar
mudando os valores do timeout*/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) //com 5 provoca timeout - aumenta pra 8 pra dar tempo
	defer cancel()
	/*Se nn recebermos a resposta dessa chamada http em 5 segundos, o contexto será cancelado por
	causa do context.WithTimeout*/
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req) //faz a requisição
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body) /*mostra no terminal - OBS: a resposta que temos aqui, esta vido da aula 3- context utilizando server http */
}
