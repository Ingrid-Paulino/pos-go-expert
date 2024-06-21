package main

import "fmt"

/*Toda vez que damos panic, travamos e finalizamos com programa*/
/*Tratativa de panico, geralmente o panic trava a aplicacao toda, vamos tratar para saber o que aconteceu e como cnverter
para n travar toda a aplicação*/
func panic1() {
	panic("Something went wrong panic1")
}

func panic2() {
	panic("Something went wrong panic2")
}

func main() {
	//recupera um panic, ao inves do programa travar, ele recupera e imprime o valor na tela e o programa n trava, pois vc vai poder tomar alguma decisão
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main: ", r) //trata o panic
			if r == "Something went wrong panic1" {
				fmt.Println("panic1 recovered")
			}

			if r == "Something went wrong panic2" {
				fmt.Println("panic2 recovered")
			}
		}
	}()

	panic1()
}
