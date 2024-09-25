package main

import (
	"log"
	"net/http"

	"github.com/Ingrid-Paulino/pos-go-expert/APIs/configs"
	_ "github.com/Ingrid-Paulino/pos-go-expert/APIs/docs" //não é usado diretamente, mas é usado pelo swagger
	"github.com/Ingrid-Paulino/pos-go-expert/APIs/infra/database"
	"github.com/Ingrid-Paulino/pos-go-expert/APIs/infra/webserver/handlers"
	"github.com/Ingrid-Paulino/pos-go-expert/APIs/internal/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
Redar a aplicação e o banco de dados em memória sqlite
1- go run main.go
2- Faça uma requisição POST para http://localhost:8080/products com o seguinte body: {"name": "Product 1", "price": 10.00} / faça pelo arquivo test/product.http
2- em outra tela, vai na raiz do projeto e rode -> sqlite3 cmd/server/test.db
3- select * from products;
*/

/*Comentarios serão lidos pelo swagger e gerarão a documentação da API - ApiKeyAuth  vai pedir para autendticar na api para poder usar, nesse caso o tokenvai estar no header com o nome de Authorization  */
/*swag init cmd/server/main.go - roda o swagger - ese comando gera a pasta docs*/

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Wesley Willians
// @contact.url    http://www.fullcycle.com.br
// @contact.email  atendimento@fullcycle.com.br

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//Forma 1
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	//fmt.Println(config.DBDriver) //config.DBDriver é posivel ser modificada config.DBDriver = "batatinha frita 123" - forma 2 não é posivel ser modificada
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	//userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JwtExpiresIn)
	userHandler := handlers.NewUserHandler(userDB, configs.JwtExpiresIn) //configs.JwtExpiresIn poderia ser passado por context igual ao configs.TokenAuth

	//roteador do framework chi - me da mais controle sobre as rotas
	r := chi.NewRouter()                                  //Cria um novo roteador chi, poderia usar o Mux tbm
	r.Use(middleware.Logger)                              //Adiciona um middleware de log //lib chi
	r.Use(middleware.WithValue("jwt", configs.TokenAuth)) //passo o valor do meu middleware de autenticação para o contexto da requisição

	//agrupa as rotas de products
	r.Route("/products", func(r chi.Router) {
		//Sempre que eu acessar qualquer URL de products, vai passar por esse middleware de autenticação
		r.Use(LogRequest)                          //criado na aula
		r.Use(jwtauth.Verifier(configs.TokenAuth)) //pega o token da requisição e joga no contexto da requisição
		r.Use(jwtauth.Authenticator)               //protege os products, quem acessar os products precisa esta autenticado no sistema e verifica  se o token se é valido //middleware de autenticação
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	//swagger: rota para acessar a documentação da API
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json"))) //http://localhost:8080/docs/index.html: rota para acessar a documentação da API

	http.ListenAndServe(":8080", r)

	//Roteador padrão do Go
	//http.HandleFunc("/products", productHandler.CreateProduct)
	//http.ListenAndServe(":8080", nil)

	//FORMA 2 do arquivo configs/config.go
	//Dessa forma nn conseguimos fazer alteração nas configuraçoes do sistema
	//config2 := configs.NewConfig()
	//fmt.Println(config2.GetDBDriver()) /*config2 não tem acesso aos campos da struct conf2, pois estão privados, então não é possivel acessar os campos da struct conf2 diretamente.
	//Mas podemos criar métodos na struct conf2 que retornam os valores dos campos da struct conf2. Dessa forma evitamos que anterem o valor da nossa config DBDriver */
}

// Middleware
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
