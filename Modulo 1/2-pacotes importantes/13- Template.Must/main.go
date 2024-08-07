package main

import (
	// "html/template"
	//ou
	"os"
	"text/template"
)

/*Qual importar "html/template" ou "text/template": eles funcionam exatamente da mesma forma, mas o html/templates sabe que a
informação vai passer passeada de html, sendo assim a aplicação vai se blindar ao maximo de alguns tipos de ataque. dica: toda vez que for
trabalhar com templates utilize html/templates caso vá exibir isso no formato html no browse. Caso for fazer isso apenas de forma textual pode utilizar o text/template */

//Go tem um sistema de template imbutido, não é necessario nenhuma biblioteca externa

type Curso struct {
	Nome         string
	CargaHoraria int
}

// rode: go run main.go
func main() {

	//OBS: mesmo exemplo que o anterio porem usando o Must. Must posibilita jogar tudo pra uma unica linha
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n"))
	//t retorna um template ou um erro caso houver
	err := t.Execute(os.Stdout, curso) //os.Stdout vai imprimir na tela(terminal) e curso são as variaveis q serão utilizadas para substituir na hora de realizar o parse
	if err != nil {
		panic(err)
	}
}
