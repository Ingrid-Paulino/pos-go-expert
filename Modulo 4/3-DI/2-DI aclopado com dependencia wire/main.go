package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// go run main.go wire.go
// ou: go run main.go wire_gen.go

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	//Injeção de dependencia
	usecase := NewUseCase(db)

	productResult, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(productResult)
}
