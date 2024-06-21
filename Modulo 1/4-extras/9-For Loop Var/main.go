package main

import "fmt"

/*
Quando rodar o codigo o retorno virar "c, c, c" é um bug que nao foi
corrigido antes pq a interpretação do pessoal go é que cada vez que for interado o loop, o v vai mudar, mas como temos
uma goroutine o escopo do for, não esta batendo no escopo da goroutine e ai ele acaba repetindo e pegando apenas o ultimo valor.
*/
// func main() {
// 	done := make(chan bool)
// 	values := []string{"a", "b", "c"}

// 	for _, v := range values {
// 		go func() {
// 			fmt.Println(v)
// 			done <- true //joga o done === true
// 		}()
// 	}

// 	for range values {
// 		<-done //libera o canal
// 	}
// }

/*Forma que o pessoal/devs corrigia isso: a comunidade redefinia o v dentro do escopo e assim a goroutine conseguia mudar*/
// func main() {
// 	done := make(chan bool)
// 	values := []string{"a", "b", "c"}

// 	for _, v := range values {
// 		v := v
// 		go func() {
// 			fmt.Println(v)
// 			done <- true //joga o done === true
// 		}()
// 	}

// 	for range values {
// 		<-done //libera o canal
// 	}
// }

// consertaram na versao 1.22.1 redefinido o v automaticamente sempre que o loop for rodado
// obs: nn vai funcionar pois nn tenho a versao 1.22.2
// func main() {
// 	done := make(chan bool)
// 	values := []string{"a", "b", "c"}

// 	for _, v := range values {
// 		go func() {
// 			fmt.Println(v)
// 			done <- true //joga o done === true
// 		}()
// 	}

// 	for range values {
// 		<-done //libera o canal
// 	}
// }

// disponivel na versao 1.22.1
func main() {

	// Forma simplificada do loop disponivel apartir da versão 1.22.1
	X := 10
	for i := range X {
		fmt.Println(i)
	}

	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true //joga o done === true
		}()
	}

	for range values {
		<-done //libera o canal
	}
}
