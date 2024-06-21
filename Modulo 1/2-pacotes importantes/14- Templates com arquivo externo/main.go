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

type Cursos []Curso

// rode: go run main.go
func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html")) //dentro do ParseFiles poderia passar um slice de strings -> []string["template.html"...]
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 10},
		{"Python", 30},
		{"C", 40},
		{"C#", 20},
	})
	if err != nil {
		panic(err)
	}
}
