package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

type Empresa struct {
	Nome string
}

// Qualquer struct que tiver o metodo Desativar esta implementando a interface Pessoa
// Interface possibilita utilizar diversos tipos de forma simples
// Qualquer struct que tiver o metodo desativar, estará implementando a interface Pessoa
// A interface do GO aceita apenas metodos, não aceita atributos/propriedades.
type Pessoa interface {
	Desativar()
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
	fmt.Println(c)
}

func (e Empresa) Desativar() {
	fmt.Printf("O nome da empresa é %s\n", e.Nome)
	fmt.Println(e)
}

// Qualquer struct que tiver o metodo desativar vai poder ser usado no parentese pessoa
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	ingrid := Client{
		Nome:  "Ingrid",
		Idade: 21,
		Ativo: true,
	}

	minhaEmpresa := Empresa{Nome: "Mercado livre"}

	Desativacao(ingrid)       //Ingrid é uma pessoa pois implementa o metodo Desativar()
	Desativacao(minhaEmpresa) //Desativação aceita minhaEmpresa, pois Empresa tem o metodo desativacao
}
