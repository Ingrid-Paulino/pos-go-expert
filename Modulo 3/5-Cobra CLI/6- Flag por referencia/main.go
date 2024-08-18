/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/Ingrid-Paulino/CLI/cmd"

// go run main.go
//go run main.go category --name=X
//go run main.go category -n=X   oh go run main.go category (tras valor default)
//go run main.go category --exists=true /go run main.go category -e=true / go run main.go category -e

func main() {
	cmd.Execute()
}
