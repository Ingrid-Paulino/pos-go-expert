package main

//QUANDO USAR OU NÃO USAR PONTEIRO:
// Não vai usar quando quiser só passar a copia dos dados para fazer uma utilização. Ex: quando eu quiser só fazer o retorno da soma de um valor
// Usar quando quiser tornar os valores passados mutaveis. Assim qualquer lugar que eu alterar o valor vai haver mudança em qualquer lugar do projeto

// resumo: vou fazer mudança na variavel, usa ponteiro. Não quero mudar o valor da variavel, passo como copia.
// Toda vez que estamos trabalhando com GO e passamos um valor de parametro, estamos fazendo uma copia do valor que esta na memoria. Não estamos passando realmente o valor que esta na mesmoria.
func soma(a, b int) int {
	return a + b
}

func soma2(a, b int) int {
	a = 50
	return a + b
}

func soma3(a, b *int) int {
	*a = 50
	return *a + *b //tenho que retornar com ponteiro, pois agr estamos mechendo com o endereço da memoria e para pegar o valor usamos o *
}

func main() {
	var1 := 10
	var2 := 20

	//Estamos passando uma coipia das variaveis var1 e var2
	println(soma(var1, var2)) //aqui os parametros estão pegando uma copia do valor das variaveis var1 e var2
	//Estamos passando uma coipia das variaveis var1 e var2
	println(soma2(var1, var2)) //mesmo alterando o valor de a dentro de soma2, var1 dentro de main continua com o memso valor, pois oq passamos pra dentro do parametro foi uma copia
	println(var1)

	var3 := 10
	var4 := 20

	//estamos passando o endereço de memoria das variaveis var3 e var4
	println(soma3(&var3, &var4)) //pasando os valores das varaveis com um & nos parametros, qualquer modificação que houver dentro da função soma3 terá modificaçao na funcao main, pois estaremos mexendo com o valor do endereço da memoria
	println(var3)                //var3 tem o valor modificado pois alteramos na func soma3
}
