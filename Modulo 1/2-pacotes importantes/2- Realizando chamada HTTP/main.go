package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//chamada http
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body) //dado com o retorno do resultado
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res)) //retorno será todo o codigo do site do google
	req.Body.Close()         //se não fechar pode vazar recursos
}
