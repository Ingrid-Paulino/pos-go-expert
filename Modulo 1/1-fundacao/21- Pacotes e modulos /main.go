package main

import (
	"fmt"

	//Para importar outros pacotes do proprio projeto em outros lugares, é necessario importar o pacote
	//Só é posivel importar outro pacote seguindo o seguinte passo:
	//Para importar outros arquivos de pastas diferentes no projeto, é necessario o arquivo go mod
	// para criar o arquivo é só rodar o comando: go mod init nomeDoMODULO(geralmente o pessoal coloca um indereço no projeto do github)
	//ex: go mod init github.com/fullcycle/curso-go
	// go mod init indica que estamos inicializando um modulo no Go
	//Logo a pós criado, rode o comando 'go mod tidy' baixa as libs externas que tiverem no arquivo criado com o comando a cima
	//O caminho aqui é o nome do modulo no arquivo go.mod + o nome do diretorio
	"github.com/Ingrid-Paulino/pos-go-expert/matematica" //tendo o arquivo go.mod eu consigo acessar os pacotes que estão em outros diretorios do projeto
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro)
	carro.Andar()

}
