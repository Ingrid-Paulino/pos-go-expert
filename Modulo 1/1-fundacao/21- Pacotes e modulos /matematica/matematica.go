package matematica

import "fmt"

// Tudo que começa com letra maiuscula é exportavel
// Para ser posivel importar funcoes, variaveis... em outros pacotes é necessario que a primeira letra co nome da func ou var esteja em maiusculo
func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() {
	fmt.Println("Carro andando")
}

// Essa func só pode ser usada dentro do arquivo matemantica.go
func soma[T int | float64](a, b T) T {
	return a + b
}

var a int = 10 //variavel a só pertence ao arquivo matematica.go

type carro struct { //struct carro só pertence ao arquivo matematica.go
	Marca string
}
