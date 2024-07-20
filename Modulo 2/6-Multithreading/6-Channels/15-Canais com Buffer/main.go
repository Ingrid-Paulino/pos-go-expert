package main

/*
	Até aqui todos os exemplos de canais que vimos, sempre os canais recebia e enviava uma mensagem de cada vez
	O Buffer é uma forma de enviar e receber mais de uma mensagem de uma vez
*/

/*
Não é recomendado ficar usando buffer em canais, pois pode causar problemas de sincronização.
Evite usar buffer em canais, a menos que você tenha um bom motivo para isso.
Dependendo da situação so vai ficar mais lento o processamento e aumentando a memoria usada

É recomendado fazer banchmark para ver se o buffer esta ajudando ou atrapalhando
Usando buffer é necessario mais threads para processar as mensagens do canal e isso pode causar problemas de sincronização e aumentar o tempo de processamento, entao é necessario fazer banchmark para ver se o buffer esta ajudando ou atrapalhando
Usando buffer é necessario mais threads para ler as mensagens e não dar o erro de deadlock, pois se o buffer estiver cheio e não tiver nenhuma thread lendo as mensagens, vai dar o erro de deadlock
*/
func main() {
	ch := make(chan int, 3) //Buffer de 2

	ch <- 1
	ch <- 2

	println(<-ch)
	println(<-ch)
}
