package math

// Modificador de acesso: maiúsculo para público, minúsculo para privado
// Se o nome da struct, função, variaveis ... for maiúsculo, ela será pública, é possivel exportar. Se for minúsculo, será privada.
// Os valores dos atributos da struct são privados, mas é possivel exportá-los usando o modificador de acesso.
type Math struct {
	A int
	B int
}

func (m Math) Add() int {
	return m.A + m.B
}

var X string = "Hello, World!"

type Math2 struct {
	a int
	b int
}

/*
Caso queira que os valores dos atributos da struct sejam privados, para ninguem ter asseso a eles diretamente, o caminho é criar funções para acessá-los.
*/
func NewMath(a, b int) Math2 {
	return Math2{a, b}
}

func (m Math2) Add2() int {
	return m.a + m.b
}

// Outra forma de fazer isso é criar uma struct privada e uma função pública para instanciá-la.
type math3 struct {
	a int
	b int
}

func NewMath2(a, b int) math3 {
	return math3{a, b}
}

func (m math3) Add3() int {
	return m.a + m.b
}
