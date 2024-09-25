package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

//Forma 1 sem Go Routine
//func main() {
//	//Primeiro roda task A e depois task B
//	task("A")
//	task("B")
//}

// Forma 2 com Go Routine
// Thread 1 é a main
// OBS: O programa não vai rodar pois a thread 1 vai finalizar antes das outras threads
//func main() {
//	//Roda de forma simultanea e concorrente
//	go task("A") //Thread 2
//	go task("B") //Thread 3
//	//Nada aqui
//	//Sair do programa
//}

// Forma 3 com Go Routine e um time sleep
// Thread 1 é a main
func main() {
	//Roda de forma simultanea e concorrente e espera 15 segundos para finalizar a thread 1 e dar tempo das outras threads rodarem
	go task("A") //Thread 2
	go task("B") //Thread 3

	//Go permite trabalhar com func anonimas
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonnymous")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}
