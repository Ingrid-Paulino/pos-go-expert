package main

import (
	"fmt"

	"github.com/ingrid-paulino/goexpert/package/math"
)

func main() {
	//Para usar p pacote math do projeto, precisamos importá-la usando o caminho relativo ao módulo ("github.com/ingrid-paulino/goexpert/package/math"). Caminho determinado pelo go.mod do projeto. Nome do pacote no arquivo go mod.
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	m.B = 4
	fmt.Println(m.A, m.B)
	fmt.Println(m)

	// A struct Math2 é privada, então não é possivel instanciá-la diretamente. Para instanciá-la, precisamos usar a função NewMath.
	m2 := math.NewMath(1, 2)
	fmt.Println(m2.Add2())
	fmt.Println(m2) //nn é possivel acessar os atributos da struct diretamente, pois eles são privados.

	fmt.Println(math.X)

	// A struct math3 é privada, então não é possivel instanciá-la diretamente. Para instanciá-la, precisamos usar a função NewMath2.
	m3 := math.NewMath2(1, 2)
	fmt.Println(m3.Add3())
	fmt.Println(m3) //nn é possivel acessar os atributos da struct diretamente, pois eles são privados.
}
