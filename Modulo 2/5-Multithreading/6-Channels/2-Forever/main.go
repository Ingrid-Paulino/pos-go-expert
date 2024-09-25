package main

// Thread 1
func main() {
	forever := make(chan bool) // Cria um canal de bool (canal vazio)

	<-forever // Esperando o canal ficar cheio para esvaziar - Recebe a mensagem do canal (bloqueante, está vazio)

	//OBS: Quando rodar o projeto,vai ficar travado nessa linha, pois o canal está vazio e não tem nenhuma goroutine enviando mensagem para ele
	//Recebemos um erro de deadlock
}
