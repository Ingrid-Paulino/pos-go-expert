package main

import "fmt"

// A linguagem go não é orientada a objetos, o que se aproxima das classes, methodos, etc é as STRUCTS
type Client struct {
	//dados que compoem a minha struct
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	ingrid := Client{
		Nome:  "Ingrid",
		Idade: 21,
		Ativo: true,
	}
	//muda valor do atributo da struct
	ingrid.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", ingrid.Nome, ingrid.Idade, ingrid.Ativo)
}
