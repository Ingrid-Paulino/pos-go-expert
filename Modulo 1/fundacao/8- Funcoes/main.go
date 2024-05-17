package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2)) //chama a função e printa

	fmt.Println(sum2(2, 2))

	fmt.Println(sum3(20, 40))

	//Funcao com dois retornos
	valor, err := sum4(10, 60)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

// 1- passa dados na entrada e recebe dados na saida
func sum(a int, b int) int {
	return a + b
}

func sum2(a, b int) int { //caso os parametros seja do mesmo tipo, podemos fazer dessa forma a tipagem
	return a + b
}

func sum3(a, b int) (int, bool) { //As funcoes podem retornar mais de um valor
	if a+b > 50 {
		return a + b, true
	}

	return a + b, false
}

func sum4(a, b int) (int, error) { //No Go nao tem exeptions como estamos acostumados(try/catch...), entt normalmente retornamos o tipo error
	if a+b > 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return a + b, nil //erro é nulo/vazio
}

// 3- passa dados de entrada e não recebe nada na saida
func sub(a int, b int) {
	if a > b {
		fmt.Println(true)
	}
	//func sem retorno
}

// 2- nao passa dados na entrada e recebe na saida
func sub2() bool {
	if 30 < 40 {
		return true
	}

	return false
}
