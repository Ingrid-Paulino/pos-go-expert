package main

import (
	"database/sql"
	"fmt"

	"github.com/Ingrid-Paulino/DI/product"
	_ "github.com/mattn/go-sqlite3"
)

// go run main.go
func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	//Injeção de dependencia
	/*
		é interesante termos uma forma de resolver esse encadeamento de dependências de forma automática, sem precisar criar manualmente cada uma das instâncias e passar como argumento para a próxima.
	*/
	repository := product.NewProductRepository(db)
	usecase := product.NewProductUseCase(repository)

	productResult, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(productResult)
}
