package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Ingrid"}`))              //recebe uma string json em um slice de bytes
	resp, err := c.Post("http://google.com", "application/json", jsonVar) //vai dar erro pq o google nn aceita um post
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil) //CopyBuffer pega os dados e joga os dados pra onde vc escolheu (os.Stdout/poderia ser outro lugar) armazenar
}
