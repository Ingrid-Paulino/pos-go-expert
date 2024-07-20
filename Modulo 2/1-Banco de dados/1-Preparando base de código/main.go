package main

import (
	"database/sql"
	"fmt"

	/*Para estabelecer concção com o banco é necessario impotar o drive do banco que vou usar*/
	_ "github.com/go-sql-driver/mysql" // _: é usado para importar o pacote sem usá-lo diretamente no código
	"github.com/google/uuid"
)

// Inicialize o pacote para poder usar o uuid: go mod init github.com/ingrid-paulino/goexprt/1.bancodedados
//go mod tidy

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

// go run server.go para executar o programa
func main() {
	//Estabelece coneccao com o banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Product 1", 19.99)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 150.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Products:")
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)") //? é um placeholder para os valores que serão inseridos. Isso evita a injeção de SQL e melhora a segurança do aplicativo. Evitando que o usuário insira comandos maliciosos no banco de dados.
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price) //Executa a instrução SQL preparada com os valores fornecidos. _ : é usado para ignorar o valor de retorno, pois não estamos usando-o.
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	//err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price) //We could use QueryRow that receives a context

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products") // Query não recebe parametros. Query retorna um conjunto de linhas que atendem à consulta fornecida. Não é necessário fechar as linhas quando terminar de usá-las, pois elas serão fechadas automaticamente quando o objeto Rows for destruído.
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id) //Exec: executa uma ação
	if err != nil {
		return err
	}
	return nil
}
