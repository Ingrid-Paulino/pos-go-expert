package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//cria arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//grava dados no arquivo e mostra o tamanho do que foi gravado
	//tamanho, err := f.WriteString("Hello, World!") //grava strings
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo")) //grava por bytes, use quando não sabe oq será gravado
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	//Abrir arquivo para leitura
	arquivo, err := os.ReadFile("arquivo.txt") //abre o arquivo e faz a leitura
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo)) //converte os bytes em string

	//Leitura de pouco em pouco, abrindo o arquivo
	/*Vamos supor que tenho 100 mega bytes de memoria e vou ler um arquivo pra fazer alguma transformação.
	O arquivo que vou ler tem 1GB, como vou ler um arquivo maior que a memoria que eu tenho?
	Conseguimos fazer isso no GO, lendo pedacinho por pedacinho do arquivo.*/
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	//bufio é um pacote do GO
	reader := bufio.NewReader(arquivo2) //conteudo bufferizado: conteudo não será lido todo de uma vez
	//de quanto em quanto o buffer vai ler?
	buffer := make([]byte, 10) //tamanho do buffer: vai ler de 10 em 10 bytes
	//loop vai ler o arquivo e pegar o valor
	for {
		n, err := reader.Read(buffer) // faz a leitura do reader baseado no tamanho que colocamos no buffer
		if err != nil {
			break
		}
		//gruda os pedaços lidos
		fmt.Println(string(buffer[:n])) //converte para string o buffer que esta sendo lido e esta juntando em un slice o conteudo que esta sedo pego
	}

	//Remove um arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
