package main

import (
	"fmt"
)

func main() {
	//exemplo 1
	// req, err := http.Get("https://www.google.com")
	// if err != nil {
	// 	panic(err)
	// }
	// /*Defer atrasa a execução do comando a pós ele, ele só roda depois que todas as linhas da funçao foram executadas*/
	// //defer é chamado por ultimo
	// defer req.Body.Close()   //se não fechar pode vazar recursos

	// res, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(res))

	//exemplo 2
	defer fmt.Println("Primeira Linha") //essa linha será executada por ultimo
	fmt.Println("Segunda Linha")
	defer fmt.Println("Terceira Linha")
	fmt.Println("Quarta Linha")

}
