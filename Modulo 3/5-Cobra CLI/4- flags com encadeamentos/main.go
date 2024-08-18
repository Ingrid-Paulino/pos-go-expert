/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/Ingrid-Paulino/CLI/cmd"

// go run main.go
// cobra-cli add create -p 'categoryCmd' -> cria o comando create sendo filho de category
// go run main.go category -> o comando create esta dentro do comando category

// cobra-cli add list -p 'categoryCmd' -> cria o comando list sendo filho de category

// go run main.go
// go run main.go category

// go run main.go category create
func main() {
	cmd.Execute()
}
