package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Ingrid-Paulino/UOW/internal/repository"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

// Testes sem usar o UOW
func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	input := InputUseCase{
		CategoryName:     "Category 1", // ID->1
		CourseName:       "Course 1",
		CourseCategoryID: 2, /*Na primeira rodada passando id 1 funciona, na segunda passando id 2 quebra, da essro de constraints
		mesmo assim salva no banco de dados a categoria, mas nn salva o curso. podemos ver o resultado dando um select no banco.
		Para tratar isso usamos o modelo de UOW que faz rollback em caso de erro, e não deixa salvar nenhum dado no banco.
		*/
	}

	ctx := context.Background()

	useCase := NewAddCourseUseCase(repository.NewCourseRepository(dbt), repository.NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
