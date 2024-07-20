package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
Como se diferenciam optimistic e pessimistic locking no contexto de banco de dados?
Lock Otimista X Pessimista:
Concorrência Thread(concorrencia):
Quando trabalhamos com thread nós estamos compartilhando memória com isso podemos
ter resultados inesperados, para resolver esse problema podemos usar lock ou seja
a segunda thread terá que esperar a execução da primeira thread para executar sua tarefa.

Concorrência Banco de Dados:
Com banco de dados podemos ter o mesmo problema que temos trabalhando com concorrência ou seja podemos
ter dois programas acessando a mesma tabela do banco de dados, ou para ser mais específico o mesmo registro
mas a primeiro programa está atualizando o registro e o segundo está lendo o registro, o que vai acontecer é
que o segundo programa vai ler um registro desatualizado por o primeiro programa já realizou uma alteração do registro.

Lock Pessimista
O lock pessimista bloqueia uma linha do banco de dados para que seja realizado uma atualização de um registro, ou seja
casa o primeiro programa esteja atualizando um registro o segundo programa não vai conseguir acessar esse registro porque
essa linha está bloqueada, isso vai gerar um custo alto porque o banco de dados vai enfileirar todas as tentativas de acesso a esse registro isso vai deixar a aplicação mais lenta.

O timeout da conexão pode derrubar essa transação, o que vai gerar um rollback da transação, isso vai fazer com que o primeiro programa perca a alteração que ele estava fazendo no registro.

Lock Otimista
No bloqueio otimista a transação não bloqueia o recurso e para manter a integridade dos dados é utilizado um campo de controle (version), se o valor do campo version do objeto for menor
do que o valor no banco de dados, significa que outra transação já alterou o registro, e portanto, o objeto está desatualizado e isto resultará em um erro.

O bloqueio otimista é indicado quando não há grandes probabilidades de concorrência entre processos. (é menos custoso que o bloqueio pessimista)
Normalmente trabalhamos com lock otimista quando não temos muita concorrência ou quando você trabalha com um tipo de banco de banco de dados que não
suporta o lock pessimista, outro ponto que devemos levar em consideração é o custo de refazer a transação
*/

/*Bloqueio pesimista e otimista no banco de dados
- Lock pesimista: Bloqueia a tabela, a linha que vc esta trabalhando no banco de dados e ninguem consegue alterar os dados. - bloqueia o registro para que nenhum outro processo possa modificá-lo até que o bloqueio seja liberado. Isso é feito usando a cláusula FOR UPDATE no SQL. ()
- lock otimista:  Versiona quando alquem faz alguma alteração no sistema.  - não bloqueia o registro, mas verifica se o registro foi modificado por outro processo antes de atualizá-lo. Isso é feito usando uma coluna de versão no registro. O GORM suporta bloqueio otimista por padrão. Para usar o bloqueio otimista, basta adicionar um campo de versão ao seu modelo e o GORM cuidará do resto.
*/

type Category struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type Product struct {
	ID         int     `gorm:"primaryKey"`
	Name       string  `gorm:"not null"`
	Price      float64 `gorm:"not null"`
	CategoryID int     `gorm:"not null"`
	Category   Category
	gorm.Model
}

func main() {
	/*
	  Antes de começar a trabalhar rode os outros projetos para criar o banco, as tabelas e os registros no banco de dados
	*/

	//Estabelece coneccao com o banco de dados
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" //parseTime=True e loc=Local são necessários para o gorm funcionar corretamente com o mysql e o go. O charset=utf8mb4 é necessário para suportar emojis. É necessário reiniciar o docker depois de adicionar esses parâmetros. Ajuda a evitar erros de timezone e de tipo de dado.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	tx := db.Begin() //inicia uma transação
	var c Category
	//bloqueia a categoria para que nenhum outro processo possa modificá-la até que o bloqueio seja liberado //Debug mostra o resultado da query no terminal
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error //Select * from categories where id = 1 for update - for update é a clausula que bloqueia a linha na tabela
	if err != nil {
		panic(err)
	}
	c.Name = "casa"
	tx.Debug().Save(&c) //update categories set name = 'Eletronicos' where id = 1
	tx.Commit()         //commita a transação
}
