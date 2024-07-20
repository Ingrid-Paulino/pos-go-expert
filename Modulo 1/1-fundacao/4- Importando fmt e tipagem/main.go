package main

import "fmt"

type ID int //E possivel criar tipos no go
var (
	j float64 = 1.2
	k ID      = 2
)

func main() {

	/*
			%v representa o valor nomeado em seu formato padrão
			%T representa o tipo do valor
			%d espera que o valor seja um tipo inteiro de base 10
			%b espera que o valor seja um tipo inteiro de base 2
			%s os bytes da string ou fatia
			%f espera que o valor tenha um tipo float

			fmt.Print(): imprime tudo o que chega ao terminal sem adicionar nenhum espaço ou novas linhas,
			a menos que seja codificado inequivocamente. fmt.Print("Hello") fmt.Print("World")

			fmt.Printf(): fornece formatação personalizada de strings de entrada usando um ou mais verbos e,
			em seguida, imprime a string formatada no terminal sem acrescentar nenhum espaço ou novas linhas
			(a menos que explicitamente codificado).  year := 2001 fmt.Printf("I was born in %d", year)

			fmt.Println(): é semelhante a fmt.Print(), a diferença é que adiciona espaços entre os argumentos e
			acrescenta uma nova linha no final. fmt.Println("Hello", "World") fmt.Println("welcome to my world!")

			fmt.Sprint(): A função formata e retorna a string de entrada sem imprimir nada no terminal.
			a := fmt.Sprint("Hello World")
	    	b := fmt.Sprint("Hello", "World")
	    	fmt.Println(a)
	    	fmt.Println(b)

			fmt.Sprintln(): é semelhante em função a fmt.Sprint(), exceto que adiciona automaticamente
			espaços entre os argumentos.
			 a := fmt.Sprintln("Hello","World")
			fmt.Println(a)

			fmt.Sprintf(): é usado para formatar uma string de entrada. Também funciona como fmt.Printf(),
			a diferença significativa é que fmt.Sprintf()retorna o valor em vez de imprimi-lo.
			name := "Chisom"
	    	s := fmt.Sprintf("My name is %s", name)
	    	fmt.Print(s)

			fmt.Scan(): coleta a entrada do usuário da entrada padrão e a armazena em argumentos sucessivos.
			Espaços ou novas linhas são considerados valores múltiplos e são armazenados em vários argumentos.
			Esta função antecipa que um endereço de cada argumento deve ser passado.
			var name string
			fmt.Println("What's your name?")
			fmt.Scan(&name)
			fmt.Println("Nice to meet you", name)
	*/
	fmt.Printf("O tipo de E é %T\n", j)  //Printf formata mensagem - %T me diz qual é o tipo de dado da variavel
	fmt.Printf("O valor de E é %v\n", j) //%v retorna o valor da variavel
	fmt.Printf("O tipo de E é %T\n", k)  //O tipo retornado aqui sera main.ID pois esse tipo foi criado no pacote main
}
