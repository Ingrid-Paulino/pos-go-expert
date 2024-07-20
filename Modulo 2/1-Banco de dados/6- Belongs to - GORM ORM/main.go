package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Vamos trabalhar com relacionamentos entre tabelas.
type Category struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

// vamos atribuir uma categoria a cada produto. Um produto pertence a uma categoria e uma categoria tem muitos produtos. Para isso, vamos adicionar uma coluna category_id na tabela products.
type Product struct {
	ID         int      `gorm:"primaryKey"`
	Name       string   `gorm:"not null"`
	Price      float64  `gorm:"not null"`
	CategoryID int      `gorm:"not null"` //adiciona a coluna category_id na tabela products
	Category   Category //Adicionamos o modelo de categoria para o ORM conseguir fazer o relacionamento
	gorm.Model          //cria alguns recursos adicionais automatico, como created_at, updated_at e deleted_at na tabela
}

func main() {

	/*
	   Antes de rodar o projeto de um drop table products; no mysql para apagar a tabela (docker) e descomente as linhas abaixo para criar a tabela novamente e registrar os produtos e categorias
	*/

	//Estabelece coneccao com o banco de dados
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" //parseTime=True e loc=Local são necessários para o gorm funcionar corretamente com o mysql e o go. O charset=utf8mb4 é necessário para suportar emojis. É necessário reiniciar o docker depois de adicionar esses parâmetros. Ajuda a evitar erros de timezone e de tipo de dado.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}) // cria as tabelas products e categories no banco de dados

	//create category
	category := Category{Name: "Category 1"}
	db.Create(&category)

	//create product
	db.Create(&Product{
		Name:       "Product 1",
		Price:      19.99,
		CategoryID: category.ID, //atribui a categoria ao produto
	})

	db.Create(&Product{
		Name:       "Product 2",
		Price:      19.99,
		CategoryID: category.ID, //atribui a categoria ao produto
	})

	//tras todos os produtos com a categoria. Se não usar o Preload, a categoria não será carregada. O Preload é usado para carregar os relacionamentos.
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}

//COMANDOS DO DOCKER
//docker-compose up -d
//docker-compose down

//docker exec -it mysql mysql -uroot -proot goexpert
//clear

//# show tables; --> mostra as tabelas
//# desc products; --> mostra a estrutura da tabela
//# select * from products; --> mostra os registros da tabela

//drop table products; --> apaga a tabela
