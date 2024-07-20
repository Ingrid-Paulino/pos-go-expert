package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int     `gorm:"primaryKey"`
	Name       string  `gorm:"not null"`
	Price      float64 `gorm:"not null"`
	gorm.Model         //cria alguns recursos adicionais automatico, como created_at, updated_at e deleted_at na tabela
}

func main() {
	/*
			Antes de rodar o projeto de um drop table products; no mysql para apagar a tabela (docker) e descomente as linhas abaixo para criar a tabela novamente
		depois atualize a linha e depois delete o registro
	*/

	//Estabelece coneccao com o banco de dados
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" //parseTime=True e loc=Local são necessários para o gorm funcionar corretamente com o mysql e o go. O charset=utf8mb4 é necessário para suportar emojis. É necessário reiniciar o docker depois de adicionar esses parâmetros. Ajuda a evitar erros de timezone e de tipo de dado.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	//Create
	db.Create(&Product{Name: "Product 1", Price: 19.99}) // ja cria o created_at, updated_at e deleted_at

	// Update
	var p Product
	db.First(&p, 1)
	p.Name = "Product 1 Updated"
	db.Save(&p) //atualiza o updated_at

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	// Delete
	db.Delete(&p2) //atualiza o deleted_at. Dessa forma, o registro não é removido do banco de dados, apenas marcado como deletado. Isso é chamado de soft delete.

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
