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

	//OBS: para rodar esse código, é necessário rodar a aula 2 e 3 para criar o banco de dados, a tabela e os registros

	// Altera um registro
	var p Product
	db.First(&p, 1)
	p.Name = "Product 1 Updated"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	// Remove um registro
	db.Delete(&p2)
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
