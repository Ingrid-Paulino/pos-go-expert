package main

import "fmt"

func main() {
	//maps: é uma strutura com key e value, não tem nececidade de ordenaçao e pode colocar o que quiser
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Wesley"])
	fmt.Println(salarios["Maria"])
	delete(salarios, "Wesley") //deleta a posicao
	fmt.Println(salarios)

	salarios["Wes"] = 5000 //Adiciona uma nova posição
	fmt.Println(salarios)

	//funcao make só cria o mapa inicial, prepara o map e vem vazio
	sal := make(map[string]int)
	// ou posso fazer assim
	sal1 := map[string]int{}
	sal1["Wesley"] = 10
	sal["Wesley"] = 10

	//Percorre o salario de todo mundo:
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	//Ignora o indice
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
