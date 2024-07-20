package main

// Thread 1
func main() {
	forever := make(chan bool) // Cria um canal de bool (canal vazio)

	//OBS: Para evitar o deadlock, vamos criar uma goroutine para segurar o processo/programa
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
	}()

	<-forever //Apos a execucao da go routine acima, teremos um deadlock, pois o canal forever esta vazio

}
