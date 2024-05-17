package main

const a = "Hello, World!" // valor de const não se altera
var (                     // Valor de var se altera
	// No GO mesmo nao passando os valores na declaracao de variavel, ele infere valores padrões
	b bool //false
	c int  //0
	/*
		Na linguagem Go, podemos declarar números inteiros com capacidades de bytes diferentes.
			rune : -2,147,483,648 ~ 2,147,483,647
			int8 : -128 ~ 127
			int16 : -32,768 ~ 32,767
			int32 : -2,147,483,648 ~ 2,147,483,647
			int64 : -9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807

			var runeVar rune = 255
			var intVar int8 = 100
			var intVar2 int16 = -200
			var intVar3 int32 = 3000
			var intVar4 int64 = -20000
	*/

	/*
		Para números inteiros, existem aqueles que possuem capacidade apenas, para valores positivos.

			byte : 0 ~ 255
			uint8 : 0 ~ 255
			uint16 : 0 ~ 65,535
			uint32 : 0 ~ 4,294,967,295
			uint64 : 0 ~ 18,446,744,073,709,551,615

			var byteVar byte = 12
			var uintVar uint8 = 35
			var uintVar2 uint16 = 114
			var uintVar3 uint32 = 1278
			var uintVar4 uint64 = 27001
	*/
	d string  //""
	e float64 //0.000000
	/*
		float32 : -3.4³⁸ ~ +3.4³⁸
		float64 : -1.7³⁰⁸ to +1.7³⁰⁸
		Também chamamos comumente também de números de ponto flutuante ou float.
		A diferença de 32 ou 64, diz respeito a capacidade de armazenamento de dados na varíavel.
		Se serão 32 ou 64 bits de capacidade.

		var varFloat32 float32 = 19.20
		var varFloat64 float64 = 61.337
	*/

	//Declarando e atribuindo variaveis
	f bool = true
	g int  = 10
)

// Declaração de multiplas variaveis
var one, two, three, four int
var one1, two1, three1, four1 = 1, 2, 3, 4

func main() {
	//variavel de escopo local
	var a2 string
	var a3 string = "X"

	//Declaracao de variavel short rain
	h := "X" //criando e atribuindo o valor na variavel
	h = "XX" //Modificando o valor da variavel

	println(a)
	println(a2)
	println(a3)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
	println(g)
	println(h)
}
