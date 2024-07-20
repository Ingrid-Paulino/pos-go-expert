package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// Configuraçoes do projeto

// ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// FORMA 1
// Variaveis de ambiente
var cfg *conf

// conf is the configuration for the application.
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"` //mapstructure: é uma tag que indica como o pacote viper deve mapear os campos da struct para as chaves do arquivo de configuração.
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"` // porta do servidor web
	JWTSecret     string `mapstructure:"JWT_SECRET"`      // chave secreta para assinatura do token
	JwtExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`   //tempo de expiração do token
	TokenAuth     *jwtauth.JWTAuth
}

// LoadConfig carrega as configurações do arquivo .env
// formatos de arquivo de configuração: .yaml, .json, .env, .toml...
// OBS: vamos usar um pacote viper para pegar as informações do arquivo .env. Daria para fazer manualmente, mas é mais fácil usar um pacote pronto.
func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config") // nome do arquivo de configuração
	viper.SetConfigType("env")        // tipo do arquivo de configuração
	viper.AddConfigPath(path)         // caminho para o arquivo de configuração
	viper.SetConfigFile(".env")       // nome do arquivo de configuração
	viper.AutomaticEnv()              // carrega as variáveis de ambiente. Se tiver uma variável de ambiente com o mesmo nome da chave do arquivo de configuração, ela vai sobrescrever o valor do arquivo de configuração.
	err := viper.ReadInConfig()       // lê o arquivo de configuração
	if err != nil {
		panic(err) // se der erro, a aplicação não vai subir
	}
	err = viper.Unmarshal(&cfg) // mapeia as configurações para a struct conf
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil) // cria uma nova instância de JWTAuth com o algoritmo de assinatura HS256 e a chave secreta. Gera tokens JWT aleatórios para cada solicitação.

	return cfg, nil
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// FORMA 2

/*
Para ter mais segurança, podemos fazer com que a struct conf seja privada e criar uma função NewConfig que retorna um ponteiro para a struct conf.
Nesse formato vamos fazer de forma que não seja posivel exportar a struct conf, ou seja, ela só vai ser visivel dentro do pacote configs.
Isso é bom porque assim garantimos que as configurações só podem ser acessadas pelo pacote configs e não por outros pacotes.
Evitando assim que as configurações sejam alteradas em tempo de execução por outros pacotes.
*/
//
//var cfg2 *conf2
//
//func NewConfig() *conf2 {
//	return cfg2
//}
//
//type conf2 struct {
//	dbdriver string `mapstructure:"DB_DRIVER"` //mapstructure: é uma tag que indica como o pacote viper deve mapear os campos da struct para as chaves do arquivo de configuração.
//	dbhost   string `mapstructure:"DB_HOST"`
//	dbport   string `mapstructure:"DB_PORT"`
//	dbuser   string `mapstructure:"DB_USER"`
//	// ...
//}
//
//func (c *conf2) GetDBDriver() string {
//	return c.dbdriver
//}
//
//// Transformaria o LoadConfig em um init
//func init() {
//	//Nesse caso o viper não consegue pegar os dados que não estao exportados na struct conf2, tem a função viper.BindEnv()(nn trenho certeza se é essa a funcao) que enxerga as variaveis não exportadas.
//	// ...
//}
