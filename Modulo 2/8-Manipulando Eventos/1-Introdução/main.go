package __Introdução

/*
EVENTOS:
- Evento é algo que aconteceu no passado.
ex:
	- pintei o carro. O evento é o carro pintado.
	- Inseri um registro no banco de dados. O evento é o registro inserido.

Podemos fazer acoes com base em eventos.
ex:
	- Inseri um novo cliente no banco
	 	- Enviar um email para o cliente
		- Enviar uma notificacao para o vendedor
		- Publicar uma mensagem em um topico de mensageria
		- Publicar uma mensagem na fila
		- Notificar um usuario no slack
*/

/*
ELEMENTOS TÁTICOS DE UM CONTEXTO DE EVENTO:
	- passo a passo para trabalhar com eventos
	 	- Evento (vai carregar dados)
	    - Operaçoes que serão executadas quando um evento é chamado (ex: quando cliente é criado, enviar email)
		- Gerenciador dos nossos eventos/operaçoes (recebe o cadastro dos eventos e dispara as operaçoes)
			- Registrar os eventos e suas operaçoes
			- Despachar / Fire no evento para suas operaçoes sejam executadas
*/
