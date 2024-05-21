package main

import "fmt"

// type assertation força o go transformar tipos

func main() {
	var minhaVar interface{} = "Ingrid Paulino"

	println(minhaVar) //retorno (0x45c580,0x47e0b8) -> retorna isso pq o GO n sabe o tipo de minhaVar

	//Para contorna isso podemos fazer o seguinte:
	println(minhaVar.(string)) //Estou afirmando para o sistema que minhaVar é do tipo string

	//Se eu tentar colocar minhaVar como um int, float64... vai dar erro, podemos fazer tratativas caso de erro:
	//Se eu nn tratar e der erro, será retornado um panic
	res, ok := minhaVar.(int) //retorna se a conversão deu certo ou não
	// se ok der false significa que deu erro a conversão e res vai dar 0
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
}
