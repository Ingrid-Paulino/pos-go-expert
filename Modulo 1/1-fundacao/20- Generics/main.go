package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

// GENERICS
// FORMA 1
func Soma[T int | float64](m map[string]T) T { //T -> indica que a func Soma pode receber int ou float
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// FORMA 2
// Podemos trabalhar com constrants tambem, que são tipos expecificos que vc cria para serem substituidos na tipagem
type Number interface {
	int | float64
}

func Soma2[T Number](m map[string]T) T { //T -> indica que a func Soma pode receber int ou float
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// FORMA 3
type MyNumber int

// Com o ~ na frente de int, o GO considera que qualquer tipagem criada como int (MyNumber), podera usar a constrant Number2
type Number2 interface {
	~int | float64 //O "~" sig que o int vai aceitar outros tipos criados com a tipagem int, como o type MyNumber
}

func Soma3[T Number2](m map[string]T) T { //T -> indica que a func Soma pode receber int ou float
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func Compara2[T any](a T, b T) bool { //Podemos tbm usar o tipo any(qualquer tipo/coisa), para fazer comparacoes com generics sempre vai dar erro, pois sempre terá a posibilidade de receber dados de tipo diferentes para comparar
	// any n funciona com comparcoes
	// if a == b { //Descomenta, estará com erro
	// 	return true
	// }
	return false
}

// Para sair do erro da func Compara2 podemos tipar com o tipo comparable, ele vai permitir comparar caso os T do parametro sejam do mesmo tipo
func Compara3[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

//DICA
//Se formos trabalhar com >,<,>=,<=... o generics vai sempre reclamar pois estaremos propicios a receber tipos diferentes para as comparaçoes
//Nesses casos podemos usar um pacote do GO chamdo constraints, que nos permite usar esses simbolos sem problema
//pacote constraints tem varios tipos já criados https://pkg.go.dev/golang.design/x/go2generics/constraints

func main() {
	m := map[string]int{"Ingrid": 1000, "Wesley": 2000, "Maria": 3000}
	m2 := map[string]float64{"Ingrid": 100.20, "Wesley": 200.30, "Maria": 300.0}
	m3 := map[string]MyNumber{"Ingrid": 1000, "Wesley": 2000, "Maria": 3000}

	println(SomaInteiro(m))
	println(SomaFloat(m2))

	//As funçoes SomaInteiro(m) e SomaFloat(m2) são identicas mudando só a tipo de dados recebido e retornado
	//Nesses casos de repetiçoes podemos usar os generics
	//GENERICS
	println(Soma(m))
	println(Soma(m2))

	//GENERICS com constrants
	println(Soma2(m))
	println(Soma2(m2))
	println(Soma3(m3))

	println(Compara(10, 10.04))
	println(Compara(10, 10))

	println(Compara3(10, 10))
	// Nao sei pq, mas mesmo os tipos a baixo sendo diferente ele esta deixando usar parametros com tipos diferentes, era pra reclamar
	println(Compara3(10.9, 10.04))
}
