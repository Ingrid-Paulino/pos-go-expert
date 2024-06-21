package main

import "fmt"

func main() {
	/*
		  Slace: No final das contas estamos trabalhando com um array debaxo dos panos
			- Tem um tamanho pra saber ate onde ir e uma capacidade pra ele saber o quanto ele consegue receber
			de informação
		    - Não tem tamanho fixo
	*/

	//adicionei valores que ele vai inicializar, poderia inicia-lo vazio
	//slice é posivel almentar o tamanho e capacidade
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) //os dois pontos (:) antes da posicao 0 significa que, tudo que estiver a direita do slice vai desaparecer (como se eu apagasse tudo que tinha escrito no meu slace)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) //deletei tudu que tinha depois da posicao 3 do meu slice, mas a capacidade continua 9 mas o tamanho 4

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // deletei tudo que tinha antes da minha posicao 1 (Ignoro as duas primeiras posicoes e pego so o restante) OBS: Se eu tirar as primeiras posicoes tesnho a capacidade diminuida

	s = append(s, 110)                                              //Aumenta a capacidade do slice
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2]) //O tamanho da capacidade dublicou. Toda vez que eu der um append no slice, caso ele n tenho capacidade, ele pega o tamanho inicial do slice e dobra a capacidade do que tinha. por tras dos panos ele cria um array maior e joga os valores dentro desse novo array.

	//slice pode ser declarado com a função make
	animes := make([]string, 3)
	animes[0] = "Gantz"
	animes[1] = "Berserk"
	animes[2] = "Attack on Titan"
	fmt.Println(animes, len(animes), cap(animes))

	animes = append(animes, "Mob Psycho")
	fmt.Println(animes, len(animes), cap(animes))

	numeros1 := []int{1, 2, 3, 4}

	fmt.Println(numeros1[0:2])
	fmt.Println(numeros1[1:4])
	fmt.Println(numeros1[0:])
}
