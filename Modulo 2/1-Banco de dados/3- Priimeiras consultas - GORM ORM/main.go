package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int     `gorm:"primaryKey"`
	Name  string  `gorm:"not null"`
	Price float64 `gorm:"not null"`
}

func main() {
	//Estabelece coneccao com o banco de dados
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//OBS: para rodar esse código, é necessário rodar a aula 2 para criar o banco de dados e a tabela

	//Select one product
	//var product Product
	//db.First(&product, 1) //SELECT * FROM products WHERE id = 1;
	//fmt.Println(product)
	//db.First(&product, 2)
	//fmt.Println(product)
	//db.First(&product, "name = ?", "Product 3") //SELECT * FROM products WHERE name = 'Product 3';
	//fmt.Println(product)

	//Select all products
	var products []Product
	db.Find(&products) //SELECT * FROM products;
	for _, product := range products {
		fmt.Println(product)
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
