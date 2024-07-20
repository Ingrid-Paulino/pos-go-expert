package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1 é a main
// Vamos trabalhar com wait groups:  uma estrutura de dados que permite que uma goroutine espere que um grupo de goroutines seja concluído.
// dessa forma a thread 1 vai esperar as threads 2, 3 e 4 terminarem para finalizar o programa
func main() {
	//waitGroup é um contador de threads
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25) //add o numero de operações que serão executadas //25 é o numero de operações que serão executadas

	go task("A", &waitGroup) //Thread 2
	go task("B", &waitGroup) //Thread 3

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonnymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait() //espera todas as operações serem finalizadas
}
