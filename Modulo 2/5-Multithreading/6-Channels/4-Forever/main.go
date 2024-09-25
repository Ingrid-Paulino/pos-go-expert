package main

// Thread 1
func main() {
	forever := make(chan bool) // Cria um canal de bool (canal vazio)

	//OBS: Para evitar o deadlock, vamos criar uma goroutine para segurar o processo/programa
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true //Para evitar o deadlock precisamos receber a mensagem do canal - Envia a mensagem para o canal (Bloqueante, Está cheio)
	}()

	<-forever // Recebe a mensagem do canal (Esvazia o canal)
}
