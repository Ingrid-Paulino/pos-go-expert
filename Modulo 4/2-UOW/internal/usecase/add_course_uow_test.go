package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Ingrid-Paulino/UOW/internal/db"
	"github.com/Ingrid-Paulino/UOW/internal/repository"
	"github.com/devfullcycle/goexpert/18-UOW/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

//testes com o UOW

//forma de rodar esse teste:
//- migrate create -ext=sql -dir=sql/migrations -seq init (se ja tiver a migrations, nn é necessario esse comando)
// docker-compose up -d
//migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
//acessar o banco de dados: docker-compose exec mysql bash --- ou --- docker-compose exec mysql mysql -u root -p courses
//- bash-4.2#: mysql -uroot -p courses
//- bash-4.2#: root
//- mysql > show tables;
//- mysql > desc courses;
//- mysql > desc courses;
//- mysql > desc categories;
//- mysql > select * from courses;

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	ctx := context.Background()
	uow := uow.NewUow(ctx, dbt)

	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	//input := InputUseCase{
	//	CategoryName:     "Category 1", // ID->1
	//	CourseName:       "Course 1",
	//	CourseCategoryID: 1,
	//}

	input := InputUseCase{
		CategoryName:     "Category 1", // ID->1
		CourseName:       "Course 1",
		CourseCategoryID: 2, // rodando a segunda vez da erro de integridade, mas como usamos uow, será feito um rolback e não vai persistir no banco nenhuma operação
		/*Na primeira rodada passando id 1 funciona, na segunda passando id 2 quebra, da essro de constraints
		mas nao salva no banco de dados a categoria como no arquivo add_course_test.go. podemos ver o resultado dando um select no banco.
		Para tratar isso usamos o modelo de UOW que fazer rollback em caso de erro, e não deixa salvar nenhum dado no banco.
		*/
	}

	useCase := NewAddCourseUseCaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
