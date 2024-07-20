package main

//Quando bildamos um arquivo/projeto, fazemos a transformaçao do projeto em um arquivo
// comando: go build main.go -> cria um arquivo main buildado (esse arquivo se transforma em binario) (arquivo aparece com o nome main)
//dando ls -lah -> conseguimos ver quantos Mega tem o arquivo
//para rodar o arquivo buildado o comando é: ./main

// é posivel escolher em qual sistema operacional vamos bildar o projeto:
// GOOS=windows go build main.go > indica que o go vai rodar na plataforma de windows (arquivo aparece com o nome main.exe)
// GOOS=linux go build main.go
// GOOS=darwin go build main.go -> para mac

// podemos buildar escolhendo a prataforma tbm
// link com todos os tipos: https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures
//GOOS=windows GOARCH=amd64 go build main.go -> GOOS: indica o sistema operaciona GOARCH: indica a plataforma
//comando 'go tool dist list' mostra a lista disponivel com o GOOS e GOARCH disponiveis
//comando "go env GOOS GOARCH" mostra qual o valor padrao do meu computador

// comando: "go env" mostra na variavel GOHOSTARCH qual é o sistema que por padrao sera buildado caso rode somente go build main.go

// OBS: caso o projeto tenha o arquivo go.mod, não preciso buildar com main.go ao final, somente dar go build. O go busca em todo o projeto o arquivo main, compila tudo e gera o arquivo e por padrao o nome do arquivo será o nome do modulo
// se eu quiser trocar o nome do arquivo buidado é so rodar o comando "go build -o nomeEscolhido"
// passos: "go mod init nome", "go build", "go build -o nomeEscolhido"
func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a < b && c > a { //e
		println("a < b && c > a")
	}

	if a > b || c > a { //ou
		println("a > b && c > a")
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}

}
