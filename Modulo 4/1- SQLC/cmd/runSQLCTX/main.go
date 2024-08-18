package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ingrid-Paulino/SQLC/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

// TX == transação
func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil) //inicia a transação
	if err != nil {
		return err
	}
	q := db.New(tx) //passa a transação para o banco
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil { //se der erro no rollback, retorna o erro do rollback e o erro original
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit() //se der tudo certo, commita a transação
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		//OBS: se der erro na criação da categoria ou do curso, sera dado um rollback e nada sera criado
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Para rodar: go run cmd/runSQLC/main.go
func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn) //Da acesso as querie/funcoes criadas da pasta internal criada pelo sqlc

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

	//courseArgs := CourseParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Go",
	//	Description: sql.NullString{String: "Go Course", Valid: true},
	//	Price:       10.95,
	//}
	//
	//categoryArgs := CategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Backend",
	//	Description: sql.NullString{String: "Backend Course", Valid: true},
	//}
	//
	//courseDB := NewCourseDB(dbConn)
	////Transaction
	////- Sera criado uma categoria e um curso, caso haja um erro entra a criacao do curso e da categoria, será dado um roolback pra desfazer tudo que foi feito.
	//err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	//if err != nil {
	//	panic(err)
	//}
}
