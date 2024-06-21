package main

import "fmt"

func main() {
	//Array: é uma variavel que tem tamanho fixo e é possivel percorrer por ele
	var myArray [3]int

	myArrayWithValues := [2]string{"blue", "yellow"} //Podemos declarar os valores na criação do array
	fmt.Println(myArrayWithValues)
	myArrayWithValues[1] = "pink"
	fmt.Println(myArrayWithValues)

	//Array sem definição de tamanho
	myArrayName := [...]string{"Clara", "Ingrid", "Heloisa", "Itamar", "Luana"}
	fmt.Println(myArrayName)

	//Atribui valores nas posicões do array
	myArray[0] = 10 //todo array sempre comeca pela posição 0
	myArray[1] = 20
	myArray[2] = 30

	fmt.Println(len(myArray) - 1)        //Retorna a ultima posicao do array
	fmt.Println(len(myArray))            //Retorna quantidade de posicoes do array
	fmt.Println(myArray[len(myArray)-1]) //Retorna ultimo valor do array
	fmt.Println(myArray[0])              //acessa a posição 0 do array
	fmt.Println(myArray[2])              //acessa a posição 2 do array

	//Percorre pelo array
	for indice, value := range myArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", indice, value)
	}
}
