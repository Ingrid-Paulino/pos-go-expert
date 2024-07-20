package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Has one -> faz uma ligação de um para um
// Vamos trabalhar com relacionamentos entre tabelas.

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int      //adiciona a coluna category_id na tabela products
	Category     Category //Adicionamos o modelo de categoria para o ORM conseguir fazer o relacionamento
	SerialNumber SerialNumber
	gorm.Model   //cria alguns recursos adicionais automatico, como created_at, updated_at e deleted_at na tabela
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
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

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{}) // cria as tabelas products e categories no banco de dados

	//create category
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	//create product
	db.Create(&Product{
		Name:       "mouse",
		Price:      19.99,
		CategoryID: 1, //atribui a categoria ao produto
	})

	//create serial number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1, //atribui o produto ao serial number
	})

	//tras todos os produtos com a categoria e o serialNumber. Se não usar o Preload, a categoria não será carregada. O Preload é usado para carregar os relacionamentos.
	/* A relação do mouse com o eletronico é que o mouse belongs to com os eletronicos. Agora o mouse tem um(has one) 123456(serialNumber). Esse serial number não pode ser mais de nenhum produto */
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number) //res: Product 1 Eletrônicos 123456

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
