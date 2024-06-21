package main

import (
	// "html/template"
	//ou
	"net/http"
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

// rode: go run main.go -- no terminal: curl localhost:8282 -- no browser localhost:8282
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 10},
			{"Python", 30},
			{"C", 40},
			{"C#", 20},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}
