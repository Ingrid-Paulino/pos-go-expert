/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/Ingrid-Paulino/CLI/cmd"

// rode o banco de dados:
// sqlite3 data.db
// .tables
//create table categories (id string, name string, description string);
// .tables

//comando: go run main.go category create -n=Cat -d=Desc //cria uma nova categoria no banco

//select * from categories;

func main() {
	cmd.Execute()
}
