package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	// compondo struct
	Endereco //é como uma herança da struct endereço
	//Cria propriedae do tipo CoresFavoritas
	FavoriteColor CoresFavoritas

	//OBS: composicao de struct é diferente de propriedade tipada
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type CoresFavoritas struct {
	Nome   string
	Existe bool
}

func main() {
	ingrid := Client{
		Nome:  "Ingrid",
		Idade: 21,
		Ativo: true,
	}

	ingrid.Cidade = "Belo Horizonte"
	//ou
	//ingrid.Endereco.Cidade = "Belo Horizonte"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", ingrid.Nome, ingrid.Idade, ingrid.Ativo)
	fmt.Println(ingrid)
}
