package main

// Thread 1
func main() {
	forever := make(chan bool) // Cria um canal de strings (canal vazio)

	//para o canal funcionar precisamos de threads/goroutines se comunicando uma com a outra
	forever <- true //Mesmo que preenchemos o canal, vamos receber erro de deadlock pois não temos um receptor (não temos uma outra thread/goroutine para receber a mensagem)

	<-forever // Recebe a mensagem do canal (Esvazia o canal)
}
