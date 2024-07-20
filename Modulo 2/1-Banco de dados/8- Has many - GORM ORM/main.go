package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Has one -> faz uma ligação de um para um (Um produto tem um serial number)
//has many -> faz uma ligação de um para muitos (Uma categoria tem muitos produtos)
//belongs to -> faz uma ligação de muitos para um (Um produto pertence a uma categoria)
// Vamos trabalhar com relacionamentos entre tabelas.

type Category struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Product []Product
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
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
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	category2 := Category{Name: "Cozinha"}
	db.Create(&category2)

	//create product
	db.Create(&Product{
		Name:       "Mouse",
		Price:      19.99,
		CategoryID: category.ID, //atribui a categoria ao produto
	})

	db.Create(&Product{
		Name:       "Panela",
		Price:      190.99,
		CategoryID: category2.ID, //atribui a categoria ao produto
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Product").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")

		for _, product := range category.Product {
			fmt.Println("- ", product.Name)
		}
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
