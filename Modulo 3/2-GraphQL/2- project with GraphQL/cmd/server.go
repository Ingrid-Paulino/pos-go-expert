package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Ingrid-Paulino/GraphQL/graph"
	"github.com/Ingrid-Paulino/GraphQL/internal/database"
	_ "github.com/mattn/go-sqlite3" //importa o driver do sqlite3 para o go abrir a conexão com o banco de dados

)

const defaultPort = "8080"

/*
Para rodar:
1 - go run cmd/server.go
2- sqlite3 data.db //para abrir o banco de dados
3- create table categories (id string, name string, description string); //para criar a tabela
4- .tables //para ver as tabelas
5- select * from categories; //para ver os registros da tabela
6- .exit //para sair do sqlite3 - nn precisa rodar esse comando
7- create table courses (id string, name string, description string, category_id string); //para criar a tabela
7 - http://localhost:8080/ //para abrir o playground do GraphQL
*/

func main() {
	db, err := sql.Open("sqlite3", "./data.db") //essa linha eu que coloquei para conectar com o banco de dados
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb, //injetei categoryDb no resolver
		CourseDB:   courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
