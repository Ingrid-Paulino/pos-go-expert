package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int     `gorm:"primaryKey"`
	Name  string  `gorm:"not null"`
	Price float64 `gorm:"not null"`
}

func main() {
	//Establishes connection to the database
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}) //This line creates the table in the database /*
	// run docker compose and then the project*/

	//Create
	db.Create(&Product{Name: "Product 1", Price: 19.99})

	//create multiple products
	products := []Product{
		{Name: "Product 2", Price: 29.99},
		{Name: "Product 3", Price: 39.99},
		{Name: "Product 4", Price: 49.99},
	}
	db.Create(&products)
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
