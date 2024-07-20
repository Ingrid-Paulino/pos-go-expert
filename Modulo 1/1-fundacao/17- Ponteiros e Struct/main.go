package main

import "fmt"

type Conta struct {
	saldo int
}

func (c Conta) simular(valor int) int { // copia
	c.saldo += valor
	return c.saldo
}

// * na struct Conta sig que agora estou mexendo com a struct do endereço da memoria, todas as alterações que houver nesse metodo, ira refletir em todo o projeto
func (c *Conta) simular2(valor int) int { // sem copia, mexendo com ponteiro
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}

	fmt.Println(conta.simular(200))
	fmt.Println(conta.saldo) //saldo continua com o valor inicial pois não fazemos alteraçoes com ponteiro

	conta2 := Conta{saldo: 100}
	fmt.Println(conta2.simular2(200))
	fmt.Println(conta2.saldo) //saldo alterado pois dentro da função simular2 estamos trabalhando com * ponteiro

}
