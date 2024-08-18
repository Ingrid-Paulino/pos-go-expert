package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ingrid-Paulino/SQLC/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// Para rodar: go run cmd/runSQLC/main.go
func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn) //Da acesso as querie/funcoes criadas da pasta internal criada pelo sqlc

	//SQLC nos tras facilidade e agilidade para manipular o banco de dados
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	result, err := queries.CreateCategoryWithResult(ctx, db.CreateCategoryWithResultParams{
		ID:          uuid.New().String(),
		Name:        "FrontEnd",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Last insert ID:", lastInsertID)
	fmt.Println("Rows affected:", rowsAffected)
	fmt.Println(result)

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	//err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	//	ID:          "88a970e6-474e-4074-a835-e360e4b4d78b ",
	//	Name:        "Backend updated",
	//	Description: sql.NullString{String: "Backend description updated", Valid: true},
	//})
	//categories, err := queries.ListCategories(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, category := range categories {
	//	println(category.ID, category.Name, category.Description.String)
	//}

	//err = queries.DeleteCategory(ctx, "88a970e6-474e-4074-a835-e360e4b4d78b")
	//if err != nil {
	//	panic(err)
	//}
	//
	//categories, err := queries.ListCategories(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, category := range categories {
	//	println(category.ID, category.Name, category.Description.String)
	//}
}
