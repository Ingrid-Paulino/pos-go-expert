package main

import (
	"fmt"

	"github.com/google/uuid"
)

//Para instalar pacotes é só dar um:
//go mod init nomeDoModulo -> cria o arquivo go.mod
//go get nomeDoPacoteLib - > instala a lib e add no arquivo go.mod
//ex: go get golang.org/x/exp/constraints -> instala a lib no projeto
//ex: go get github.com/google/uuid
//podemos rodar "go mod tidy" -> Ele baixa os pacotes que não foram baixados e remove os pacotes que não estão sendo usados

func main() {
	fmt.Println(uuid.New())
}
