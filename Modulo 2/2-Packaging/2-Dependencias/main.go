package main

import (
	"fmt"

	"github.com/google/uuid"
)

// CIMANDO: go mod tidy - baixa as dependencias do projeto, atualiza as dependencias para as versões mais recentes ou remove as dependencias que nn estão sendo usadas no projeto.
// go.sum - arquivo de controle de versões das dependencias do projeto. Mantem as versaoes e não fica alterando toda vez que eu rodo o go mod tidy.
func main() {
	fmt.Println(uuid.New().String())
}
