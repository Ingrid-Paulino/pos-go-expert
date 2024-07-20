package __Iniciando

//Iniciando com processos
/*
Processos:
	Os processos são programas em execução.
	Nas nossas aplicaçoes temos processos que ficam rodando.
	Os processos ficam dentro de caixas e são contidos pelo sistema operacional.
	Os processos são isolados/independentes e não compartilham memória.
*/

/*
Threads:
	As threads são como processos, mas elas compartilham memória.
	threads ficam dentro do processo.
	As threads são executadas dentro de um processo.
	As threads são mais leves que os processos.
	As threads são mais rápidas que os processos.
	As threads são mais eficientes que os processos.
	As threads são mais escaláveis que os processos.
	Permite dentro do processo, trabalhar de forma paralela e concorrente.
*/

//Introdução a concorrência e Mutex
/*
 As threads compartilham memomória, temos o problema de race condition quando compartilhamos memória. Para resolver esse problema, podemos usar mutexes.
 race condition é quando duas ou mais threads tentam acessar e manipular a mesma variável ao mesmo tempo, e pelo menos uma delas tenta modificar a variável.
 Da conflito na memoria e pode gerar um resultado inesperado.

 Mutex: é um mecanismo de sincronização que permite que apenas uma thread possa acessar uma variável compartilhada por vez.
 Mutex é um Lock, que é uma trava que impede que outras threads acessem a variável compartilhada enquanto uma thread está usando a variável.
 Mutex faz Unlock, que é destravar a variável compartilhada para que outras threads possam acessá-la. Dessa forma evitamos a concorrência de threads.
*/

/*
	Concorrencia: é quando duas ou mais threads estão em execução ao mesmo tempo.
*/

//Concorrência Vs Paralelismo Vs Go
/*
	O GO trabalha com concorrência e não com paralelismo. O paralelismo é quando duas ou mais threads estão em execução ao mesmo tempo em diferentes núcleos de CPU.
	No GO temos concorrencia que ativa a necessidade de paralelismo quando necessário.
	Concorrencia é quando duas ou mais threads estão em execução ao mesmo tempo, mas não necessariamente em diferentes núcleos de CPU.
	Concorrencia não é necessario esperar o fim de uma thread para começar a outra.
	Com Concorrencia temos varias threads em execução ao mesmo tempo, mas não necessariamente em diferentes núcleos de CPU.

	A CPU é o processador do computador, é o cérebro do computador, é onde os cálculos são feitos.
	A CPU posssibilita trabalhar com o paralelismo, que é quando duas ou mais threads estão em execução ao mesmo tempo em diferentes núcleos de CPU.
	Se o computador tiver apenas um núcleo de CPU, não será possível trabalhar com paralelismo, pois só será possível executar uma thread por vez.
	Nesse caso, as threads serão executadas de forma concorrente, mas não em paralelo.
	Se o computador tiver mais de um núcleo de CPU, será possível trabalhar com paralelismo, pois será possível executar várias threads ao mesmo tempo, uma em cada núcleo de CPU.
*/

//Multithreading
/*
Multithreading:
	É quando um programa tem várias threads em execução ao mesmo tempo.
	Usar multithreading é caro, pois cada thread tem seu próprio contexto de execução, o que consome recursos do sistema.
	Mas o GO é uma linguagem que permite trabalhar com multithreading de forma eficiente, com mais liberdade.
	Usar multithreading é uma forma de melhorar o desempenho de um programa.
*/

//Schedulers
/*
Scheduler:
	É o responsável por gerenciar as threads.
	É o responsável por decidir qual thread será executada em qual núcleo de CPU.
	É o responsável por decidir quando uma thread será executada e quando será pausada.
	É o responsável por decidir quando uma thread será interrompida e quando será retomada.
	É o responsável por decidir quando uma thread será finalizada e quando será reiniciada.
	É o responsável por decidir quando uma thread será escalonada e quando será desescalada.
	É o responsável por decidir quando uma thread será bloqueada e quando será desbloqueada.

	Temos dois tipos de schedulers:
		- Preemptive: é o scheduler que decide quando uma thread será interrompida e quando será retomada. (Coloca um tempo determinado para cada thread/tarefa)
          Interrompe uma thread para dar a vez a outra. A vantagem é que garante que todas as threads tenham a oportunidade de serem executadas, mas a desvantagem é que pode haver atrasos devido à troca de contexto.
          Garantimos que não vamos ficar travados em uma thread, pois o scheduler vai interromper a thread e dar a vez a outra.

		- Cooperative: é o scheduler que decide quando uma thread será pausada e quando será retomada. (A thread decide quando vai pausar)
		Espera uma tarefa terminar para dar a vez a outra. A vantagem é que não há atrasos devido à troca de contexto, mas a desvantagem é que se uma thread ficar em um loop infinito, as outras threads não terão a oportunidade de serem executadas.

OBS: No Go não é usado o scheduler padrão do sistema operacional, mas sim um scheduler próprio do Go, que é mais eficiente e permite trabalhar com multithreading de forma eficiente.
*/

//OBS: Tudo falado a  cima, é como o sistema operacional e outras linguagens trabalha.

//Go e suas green threads - (Forma que a linguagem GO trabalha com multithreading)
/*
	A linguagem GO tem um RUNTIME(Motor para a linguagem funcionar - todo o codigo que permite a linguagem trabalhar)
	Quando pedimos pra criar nova threads, para trabalharmos com multithreads no GO, o GO não vai ate o sistema operacional pra gerar uma nova thread.
	No Runtime do GO, ele tem um proprio gerenciamento de threads. Trabalhando com Green threads, são como thrads falsas.

	Quando pedimos pro go para criar uma nova thread, ele não vai ate o sistema operacional e gera. Ele gera uma thread falsa, onde ele mesmo consegue gerenciar as tarefas que essas threads estão trabalhando.
	Sendo assim, o GO tem seu proprio Scheduler que trabalha de forma cooperativa.

	é mais barato trabalhar com threads no GO pois usamos uma quantidade baixissima de memoria

*/
/*
	Green threads:
	As green threads são threads leves, que são criadas e gerenciadas pelo próprio programa, e não pelo sistema operacional.
	As green threads são mais rápidas que as threads normais, pois não precisam fazer chamadas de sistema para o sistema operacional.
*/

/*
Go Routine:
	É uma função que é executada de forma assíncrona, concorrente, paralela, independente, isolada.
*/

/*
WaitGroup:
	É uma estrutura de dados que permite que uma goroutine espere que um grupo de goroutines seja concluído.
	Separada em 3 partes:
		- Add qtd de tarefas/operacoes que vamos executar
		- Informa que voce terminou uma operação
		- Espera todas as operações terminarem

*/

/*
6-Channels:
	É uma estrutura de dados que permite a comunicação entre goroutines.
	É uma estrutura de dados que permite a sincronização entre goroutines.
	É uma estrutura de dados que permite a passagem de mensagens entre goroutines.
	É uma estrutura de dados que permite a troca de dados entre goroutines.
	É uma estrutura de dados que permite a coordenação entre goroutines.
    o canal sempre vai receber um dado ou enviar um dado

	Nao é posivel fazer uma leitura e escrita ao mesmo tempo no channel, pois ele é bloqueante. (Se uma goroutine estiver lendo, a outra goroutine vai ter que esperar a leitura terminar para poder escrever)

	Temas tratados na pasta de channels:
		- Channels
		- Forever: problema de deadlock
		- Range
		- Select
		- Workers
		- WaitGroup
		- Direction
		- Load Balancer
		- Buffer Channels

*/
