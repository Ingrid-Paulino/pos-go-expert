package main

//Isso é uma gambiarra para o pacote achar o caminho correto localmente. na aula 4 será ensinado a forma correta de importar um pacote localmente.
import "github.com/ingrid-paulino/goexpert/package/math" //Importando o pacote math, mas se não estiver publicado no github, nn é possivel importá-lo. para isso, precisamos usar o go mod replace.

// go mod replace - substitui o pacote math do github pelo pacote math do projeto.
// go mod edit -replace github.com/ingrid-paulino/goexpert/package/math=../math - substitui o pacote math do github pelo pacote math do projeto.
// go mod tidy
//URL relativa para no go.mod --> isso suja o go.mod, use sempre workspaces para evitar isso. Aula 4.

func main() {
	m := math.NewMath(1, 2)
	println(m.Add2())
}
