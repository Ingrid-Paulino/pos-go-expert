/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/Ingrid-Paulino/CLI/cmd"

// Quando rodar o comando go run main.go ele mostra todas as opções de comandos que criei na pasta cmd
// Se rodar o comando go run main.go ping ele executa o comando ping
// O comando completion gera um arquivo de autocompletar para o shell que estiver usando
func main() {
	cmd.Execute()
}
