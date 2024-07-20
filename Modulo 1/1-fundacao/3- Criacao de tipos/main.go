package main

import (
	"fmt"
	"reflect"
)

type ID int //E possivel criar tipos no go

var (
	i ID = 1
)

func main() {
	println(i)
	fmt.Println(reflect.TypeOf(i)) // tipo main.ID -> tipo ID do arquivo main
}
