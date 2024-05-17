package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	FavoriteColor CoresFavoritas
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

// struct tem metodos
func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
	fmt.Println(c)
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

	ingrid.Desativar() //chamando o metodo //A mudança só fica dentro do metodo. Pra mudar permanente o valor do objeto ingrid teriamos que usar ponteiro

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", ingrid.Nome, ingrid.Idade, ingrid.Ativo)
	fmt.Println(ingrid)
}
