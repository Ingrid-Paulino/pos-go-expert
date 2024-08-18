package main

import (
	"fmt"

	//eventsPrivado "github.com/Ingrid-Paulino/fcutils-secret/pkg/events" //repositorio privado: não é posivel importar o pacote como fazemos normalmente com os pacotes publicos
	"github.com/devfullcycle/fcutils/pkg/events" //repositorio publico: consigo importar o pacote sem problemas
)

// Opçoes para importar pacotes privados sem problemas
//OBS: NÃO CONSEGUI FAZER ESSA AULA POIS PRECISAVA DE ALGUMA AUTENTICACOES PARA CONFIGURAR O PACOTE PRIVADO E PREFERI NÃO FAZER POR ESTAR NO PC DA EMPRESA
/*
se der "go env" no terminal, vai aparecer uma lista de variaveis de ambiente
se der "go env | grep GOPRIVATE" no terminal, vai aparecer a variavel de ambiente GOPRIVATE
   - Se quisermos baixar pacotes privados, precisamos adicionar o valor do pacote privado na variavel de ambiente GOPRIVATE (ex: GOPRIVATE=github.com/Ingrid-Paulino/*)
	-  GOPRIVATE=github.com/ -> Se colocar só isso va variavel, ela vai considerar que o github inteiro é privado
	-  GOPRIVATE=github.com/Ingrid-Paulino/* -> Se colocar só isso va variavel, ela vai considerar que o github Ingrid-Paulino inteiro é privado
	- Coloque sempre o nome da parata privada GOPRIVATE=github.com/Ingrid-Paulino/jcutils-secret,github.com/fulanotal/outropacote -> com virgula podemos adicionar varios


*/

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
	//ed2 := eventsPrivado.NewEventDispatcher()
	//fmt.Println(ed2)

}
