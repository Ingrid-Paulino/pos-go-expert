// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	"github.com/fsnotify/fsnotify"
// )

// /*vamos fazer uma forma de recarregar variaveis de ambiente com uma nova senha
// por exemplo, sem ter que parar a minha aplicação pra subir de novo

// O fsnotify vai nos ajudar a ficar lendo nossas variaveis de ambiente no arquivo .env e se ele for alterado
// ele vai recaregar no nosso sistema
// */

// type DBConfig struct {
// 	DB       string `json:"db"`
// 	Host     string `json:"host"`
// 	User     string `json:"user"`
// 	Password string `json:"password"`
// }

// var config DBConfig

// func main() {
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer watcher.Close()
// 	MarchalConfig("config.json") //se esse arquivo for modficado, ele sera atualizado automaticamente sem ter que parar a aplicação/programa
// 	fmt.Println(config)

// 	/*Vamos ficar rodando uma thread verificando eventos do watcher, pra ver se o arquivo config.json
// 	foi modifivado */
// 	//nn precisamos derubar a aplicacao pra todar tudo de novo
// 	done := make(chan bool)
// 	/*OBS: com esse sistema toda vez que eu mudar algum valor do arquivo config.json, vai atualizar o programa sem
// 	ter que parar, salvar e subir de novo*/
// 	go func() {
// 		/*
// 			Por mais que estamos trabalhando com um loop, a aplicação vai ficar parada no select, ate ela receber um evento
// 			de watcher que vem do Sistema operacional. Ela não vai ficar consumindo o tempo todo.
// 		*/
// 		for {
// 			select { //select espera alguma coisa chegar no chan (canal)
// 			case event, ok := <-watcher.Events: //fica escutando se tem algum evento em watcher
// 				if !ok {
// 					return
// 				}
// 				fmt.Println("event: ", event)
// 				//Posso fazer condiçoes para saber qual evento aconteceu
// 				if event.Op&fsnotify.Write == fsnotify.Write { //fsnotify tem varias funcoes, a write verifica se ouve alguma escrita
// 					MarchalConfig("config.json")
// 					fmt.Println("modified file: ", event.Name)
// 					fmt.Println(config)
// 				}
// 			case err, ok := <-watcher.Errors:
// 				if !ok {
// 					return
// 				}
// 				fmt.Println("error: ", err)
// 			}
// 		}
// 	}()
// 	err = watcher.Add("config.json") //watcher fica escutando o meu arquivo
// 	if err != nil {
// 		panic(err)
// 	}
// 	<-done //o canl done esta segura o programa

// }

// func MarchalConfig(file string) {
// 	data, err := os.ReadFile(file)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = json.Unmarshal(data, &config)
// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	MarshalConfig("config.json")

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event :", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig("config.json")
					fmt.Println("modified file:", event.Name)
					fmt.Println(config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()
	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}
	<-done
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}

//todo testAR RODAR POIS NN FUNIONOU
