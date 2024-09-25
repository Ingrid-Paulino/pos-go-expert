# PÓS-GO-EXPERT

## MÓDULO 1

## MÓDULO 2

## MÓDULO 3

## MÓDULO 4

## GO INTERNALS

Curso de Go Internals, que aborda os detalhes internos da linguagem Go, como o garbage collector, o scheduler, a memória,
a stack, a heap, entre outros. O curso é ministrado por William Kennedy, que é um dos maiores especialistas em Go do mundo.
Abordaremos fundamentos de computação em geral, como memória, processos, threads, heap, garbage collector, entre outros.

### PROCESSOS

O computador consegue executar multiplas tarefas pois ele tem multiplos nucleos de processamento (CPU multicore). Sendo possivel executar
varios processos em paralelo(ao mesmo tempo). Cada processo tem seu proprio espaço de memoria, e cada processo é isolado do outro.

Analisar por dentro de um processo pode fazer diferença no desempenho do programa.

#### O que é um processo?

- Instancia de um programa em execução.
- Componentes
  - Endereçamento (região da memória dedicada ao processo)
  - Contextos
    - Conjunto de dados que o SO "salva" para gerenciar um processo.
    - Ex: Program Counter(PC) ou Instruction Pointer(IP), que é um registrador que aponta para a proxima instrução a ser executada.
      - Possui o endereço da próxima instrução que o processador irá executar.
      - Auxilia no Context Switch, que é a troca de contexto entre processos.
  - Registros de Processador
    - Áreas que temporariamente armazenam no CPU dados e endereços para realizar a execução.
    - Dados
      - Ex: Realiza operaçoes aritmeticas e lógicas.
    - Registro de Endereço
      - Armazenamento em memória, incluindo stack pointers
      - Ex: Ao acessar uma variável o CPU possui em registro na mem;oria para guardar seu valor.
    - Conceito: Heap vs Stack
      - Heap
        - Utilizado para alocação de memória dinâmica. Cresce e encolhe em tempo de execução conforme a necessidade de mais ou menos espaço.
      - Stack
        - Armazena informações de controle para chamadas de função, como endereços de retorno, e parâmetros de função.
          Segue uma estrutura LIFO (Last In, First Out), onde o último item a ser inserido é o primeiro a ser removido.
    - Registro de Status/Flags
      - Fornecem os status recente das operações realizadas pelo CPU
      - Trabalha através de bits específicos (flags)
      - Ex:
        - Flag Zero (Z): Resultado de uma operação o qual o resultado é zero. Decide o fluxo do programa baseado nesse valor.
        - Flag Signal (S) ou Negative (N): Indica o resultado de uma operação positiva ou negativa.
        - Flag Overflow: Produz resultado além da capacidade do registrador.

#### Ciclo de vida de um processo

- Criação:
  - Um novo processo é criado quando um programa solicita a execução de um processo, por meio de chamadas de sistema como fork() no UNIX/Linux ou CreateProcess() no Windows
- Execução:
  - O processo está ativamente sendo executado pela CPU. Pode alterar entre os estados de "executando" e "pronto" (para ser executado)
  - Waiting/Blocked:
    - O processo é suspenso e colocado em espera até que um evento externo ócorra. Comum em operações de
- Termination: O processo completa sua execução ou é forçadamente terminado.
  I/O, onde o processo aguarda pelo término de uma leitura de disco ou recebimento de uma entrada de rede.
  - Exit: Conclusão bem-sucedida do processo após completar suas instruções.
  - Killed: Interrupção por um erro de execução ou por ser terminado por outro processo (por exemplo, através do comando "kill").

#### Criação de um novo Processo no SO

- UNIX/Linux:
  - fork()
  - Clona o processo atual
  - Gerado um processo filho
  - fork() retorna um valor diferente para o processo pai (PID)
  - Processo pai e filho são quase idênticos, porém os valores na memória são copiados para outro
    endereçamento separado e independente
  - Processo pai recebe um PID (valor inteiro positivo) do filho quando o fork() é chamado.
  - Processo filho retoma o PID 0-Indicando que é um processo filho.
  - Agora que o fork() foi realizado, cada processo pode seguir para caminhos diferentes.

#### Gerenciamento de processos

- Scheduler

  - Decide qual processo será executado
  - Alterna entre processos
  - Possui diversos algoritmos para tentar maximizar o uso da CPU
  - Scheduler pode:
    - Selecionar processos de uma fila que estão "ready queue"
    - Alocar CPU: Mudança de estado: Ready to Running
    - Retirar CPU: I/O, etc.

- 2 Tipos de Schedulers:

  - Colaborativo/Cooperativo
    - Processos que estão sendo executados tem controle quando liberam a CPU para outros
      processos.
  - Preemptivo

    - SO tem a capacidade de interromper um processo em execução e ceder o uso da CPU para outro processo. Trabalha de forma mais "justa".

  - Pontos a serem levados em consideração:
    - Processos cooperativos: Processos podem monopolizar a CPU
    - Processos preemptivos: Muitas mudanças de contexto (context switching)

### THREADS

- Treads são unidades básicas de utilização de CPU que fazem parte dos processos. (1 processo pode ter N threads)
- Threads são sequências de execução dentro do mesmo processo, compartilhando o mesmo espaço de memória e recursos.
- Dentro de um único processo, várias threads podem existir, cada uma executando diferentes
  partes do programa.
- Paralelismo vs Concorrência: Com múltiplos CPUs conseguimos atingir paralelismo. Com apenas um núcleo, trabalhamos de
  forma concorrente (simulando paralelismo).

OBS 1: O processo A não consegue acessar a memória do processo B, pois cada processo tem seu próprio espaço de memória.
Um não mexe no espaço de memória do outro. Quando temos threads rodando dentro de um processo, elas compartilham o mesmo espaço de memória.
Aendo assim, posso ter a thead 1 e a thread 2 dentro do meu processo, tanto a tread 1 e 2 podem acessar os mesmos valores na memoria e
poder fazer modificações. Porém, isso pode gerar problemas, pois se as threads estiverem acessando o mesmo valor na memória, pode ocorrer condição de corrida
entre elas, onde uma thread pode sobrescrever o valor da outra, gerando um comportamento indesejado.
Para evitar isso, é necessário utilizar mecanismos de sincronização, como mutexes, semáforos, etc.

OBS 2: Com múltiplos CPUs conseguimos atingir paralelismo, pois posso ter um processo e varias threads rodando em paralelo, cada uma em um CPU diferente.
Se eu mudar pra um unico CPU, as threads vão rodar de forma concorrente, simulando o paralelismo.
Conseguimos atingir paralelismo quando trabalhamos com múltiplos core (multiplos CPUs), e com apenas um núcleo, trabalhamos de forma concorrente (simulando paralelismo).

#### Custo de memória

- Threads, obviamente, ocupam menos espaço na memória do que um processo, pois elas compartilham a mesma memória do processo.
- Cada thread possui sua stack independente e isolada.
- Cada thread ocupa ~= 2MB (linux) (~= -> aproximadamente)

OBS: As treads do GO, ocupam 2KB de memória ao enves de ocupar 2MB, pois o GO possui um gerenciador de memória proprio, que é mais eficiente que o gerenciador de memória do SO.
Sendo assim, o GO consegue gerenciar melhor a memória, e consegue alocar menos memória para as threads.

### ARQUITETURA DO RUNTIME

Runtime do Go é toda base do codigo que é necessario para executar um programa em Go. Parte do Runtime: redes, scheduler, garbage collector(responsavel por fazer
a a limpeza de memoria, liberar memoria que não está sendo utilizada), Stack Management, Network Poller, Reflection, goroutines(threads gerenciadas pelo Runtime, são threads mais leves),
channels(possibilita trabalhar de forma concorrente, faz comunicação e sincronização entre threads), memoria alocada, etc.
O runtime do Go é o coração da linguagem, é o que faz a linguagem funcionar.

![Captura de Tela 2024-09-18 às 07.24.29.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_uwJGeY%2FCaptura%20de%20Tela%202024-09-18%20%C3%A0s%2007.24.29.png)

#### Padrão M:N

- Threads Virtuais vs Threads Reais
  - Threads Virtuais (M): São as threads gerenciadas pelo runtime do Go. São mais leves que as threads reais, pois o runtime do Go
    consegue gerenciar melhor a memória. Podem ser chamadas de: user land, green threads, light threads.
  - Threads Reais (N): São as threads do SO, que são as threads que o SO consegue gerenciar.
- Modelo de agendamento de tarefas

  - M tarefas são mapeadas para N threads: Posso pegar diversas tarefas que quero fazer um agendamento e atribuilas
    para as threads virtuais realizar, e as threads virtuais vão ser executadas pelas threads reais. É possivel ter uma thread real
    que tenha 4 ou mais threads virtuais rodando dentro dela.

  Esse tipo de arquitetura é interesante pois o GO consegue economizar memoria e assim não ficamos chamando o SO toda hora para criar uma nova thread.

#### Goroutines

- Funções / Métodos que são executadas de forma concorrente.
- São threads gerenciadas pelo runtime do Go.
- Muito mais barata do que criar novas threads no SO(2kb).
- Mais rápido criar e destruir.
- Compartilham os mesmos endereços de memória do programa em Go. Possuem Stacks independentes.

#### M:P:G

Modelo que o GO utiliza para gerenciar os programas em Go:
![Captura de Tela 2024-09-18 às 07.15.48.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_BpOh5g%2FCaptura%20de%20Tela%202024-09-18%20%C3%A0s%2007.15.48.png)

- Go Runtime: Gerencia tudo executado pela linguagem.
- Machine: Cria no SO threads reais.
- Processor: Responsavel por executar as goroutines.
  Conforme a linguagem sente a necessidade, ela cria mais threads reai, processadores, e conforme a necessidade, ela cria mais threads virtuais.

#### GoMAXPROCS

- runtime.GOMAXPROCS() -> é uma função do runtime, onde é possivel definir o número de processadores que o programa pode utilizar.
- Go cria um P (Processor) por Núcleo Computacional.
- Go tende a criar um M (Machine - Treads) para atribuir para cada P".
- O valor de uma Machine por Processor não é fixo.
- Go pode criar mais threads no SO se as atuais estiverem bloqueadas po I/O ou outro motivo de executar as Goroutines.
- Objetivo é sempre manter os "Ps" ocupados, sem tempo ocioso.

#### Scheduler

- Responsável por agendar as goroutines nos processadores.
- O Go tem um scheduler próprio, que é diferente do scheduler do SO.
- Gestão de como e quando as tarefas são executadas em threads do sistema operacional
- Decide qual tarefa deve ser executada em qual thread e em que momento
- Gerencia o balanceamento de carga entre diferentes threads ou processadores lógicos, garantindo que nenhuma thread fique sobrecarregada enquanto outras estão ociosas
- Gerencia questões como sincronização, mutex, racing conditions, deadlocks, etc
- Scheduler faz parte do Runtime. Trabalha de forma adaptativa.
  - Atribuição de tarefas
  - Balanceamento de carga
  - Gerenciamento de concorrência
  - Trabalha de forma não cooperativa com preempção (versão >= 1.14)
- Scheduler determina o estado de cada Goroutine
  - Runnung (executando)
  - Runnable (Fila) - Não está executando, mas está pronta para executar
  - Not runnable (bloqueada fazendo I/O por exemplo) - Ja executou, não finalizou e está esperando algo
  - Work Stealing (Roubo de trabalho)
    - Se o P está ocioso (Idle)
    - Ele rouba Goroutines de outro P ou mesmo da fila global de Goroutines
    - Verifica 1/61 do tempo, evitando overhead para evitar buscar a na fila global o tempo todo

![Captura de Tela 2024-09-19 às 06.17.47.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_3dmiTV%2FCaptura%20de%20Tela%202024-09-19%20%C3%A0s%2006.17.47.png)

- Processors (P) pegam as Goroutines da fila global e executam de forma local, chamando o Machine (M) para executar cada Goroutine em uma thread do SO.

#### Preempção vs Cooperação/Colaboração

Preempção: É quando uma goroutine é executada por um tempo muito longo, e o scheduler decide interrompê-la para dar a vez a outra goroutine.
Depois de um tempo, a goroutine interrompida pode ser retomada. (a partir da versão 1.14)

Cooperação: Espera terminar a execução de uma goroutine para dar a vez a outra goroutine. Roda até o fim, sem interrupções do scheduler. (versão < 1.14)

#### Schedtrace (Debug) -> Debuga o scheduler do GO funcionando

LIB PARA TESTAR CHAMADAS DE CONCORRENTES:
go-wrk: https://github.com/tsliwowicz/go-wrk - (tem que instalar no pc)

- go-wrk é uma ferramenta moderna de benchmarking HTTP capaz de gerar carga significativa quando executada em uma única CPU multi-core.
  Ela se baseia em rotinas e agendadores da linguagem go para IO assíncrono e simultaneidade nos bastidores.

comando:
1: GOMAXPROCS=1 GODEBUG=schedtrace=1 go run main.go

- GOMAXPROCS=1 -> é uma variavel de ambiente que é possivel definir o número de processadores que o programa pode utilizar.
- GODEBUG=schedtrace=1 -> Debuga o scheduler do GO funcionando
  2: go-wrk -c 20 -d 100000 http://localhost:6666/leak - (20 é a qtd de pessoas que vao acessar(qtd de requisicoes), 100000 é a duração do tempo que vai rodar, http://localhost:6666/leak é a url que vai ser testada)

### GERENCIAMENTO DE MEMÓRIA NO SO

Temos dois tipos de memória:

- Memória de rápido acesso

  - L1 - 64kb
  - L2 - 0.5mb
  - L3 - 8mb
  - Esse tipo de me,ória é utilizada como cache pelo CPU
  - Fica no mesmo chip do CPU
  - Esses layers de memoria ficam soldados no CPU, e são muito mais rápidos que a memoria RAM. O CPU conseguem acessar essas memorias
    rapidamente(nanosegundos). Quem gerencia essa memoria é o CPU.

- Memória de acesso lento
  - Pentes de memória convemcionais, memoria que colocamos na placa mãe do computador.
  - É ligada através de um barramento (canais de comunicação entre a CPU e a memória)
  - Endereço de memória são referenciados em forma hexadecimal(0-9 A-F).

#### Funcionamento da memória e Stack

- Threads, obviamente, ocupam menos espaço na memória do que um processo, pois elas compartilham a mesma memória do processo.
- Cada thread possui sua stack independente e isolada. (Stack são memorias que são muito barata e rapidas de acessar)
- Cada thread ocupa ~= 2MB (linux) (~= -> aproximadamente)

![Captura de Tela 2024-09-19 às 07.12.29.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_v3p44M%2FCaptura%20de%20Tela%202024-09-19%20%C3%A0s%2007.12.29.png)

- Stack: Quem gerencia a stack é o CPU. A stack é uma memoria que é muito barata e rapida de acessar. A stack é utilizada para guardar valores temporarios, como retorno de uma função, parametros de uma função, variaveis etc.
- Heap: Quem gerencia são os desenvolvedores.
  - Alocação de memória dinâmica
  - Acessível globalmente
  - Reutilizavel
  - Gerenciada pelo programador
  - Suporta estruturas de dados complexas
  - Leaks de memoria (quando o programa começa a locar memoria e não libera, gerando um consumo excessivo de memoria)
  - permite trabalhar com concorrencia
  - Mais lento que a stack

#### Gerencia de memória no GO

- Go possui um gerenciador de memória proprio, que é mais eficiente que o gerenciador de memória do SO.
- Utiliza como base o TCMalloc (desenvolvido pela google)
  - Ao longo do tempo. o alocador tomou diferentes caminhosdo TCMaloc
  - O proprio runtime do GO é responsavel por trabalhar com a alocação de memoria.
- Nome do alocador é "mallocgc"

#### Garbage Collector

- Responsável por gerenciar a memória alocada e desalocada.

O Garbage Collector (GC) é um mecanismo automático de gerenciamento de memória que busca, identifica e libera memória
que não está mais sendo utilizada pelo programa. Isso é crucial para prevenir vazamentos de memória e garantir a eficiência
do uso de recursos de memória

O GC entende a memoria como um conjunto de objetos, e não como um conjunto de bytes. Ele consegue identificar quais objetos
estão sendo utilizados fazendo marcaçoes nos objetos que estão sendo utilizados, e os objetos que não estão sendo utilizados.

EX: Rastreia os objetos(variaveis, structs...) utilizados de preto e os objetos não utilizados de branco. Depois vare os objetos brancos da memoria.

#### GOGC Variável

- GC Percentage (GOGC -> garbage collector do GO)
  - é uma variavel de ambiente em porcentagem que é possivel definir o valor do garbage collector do GO.
  - Define o tamanho da heap quando o GC dever ser acionado/chamado.
  - Padrão é 100%. Se não tiver nenhuma variavel de ambiente definida, o GO vai utilizar o valor padrão que é 100%.
  - Ex: Se o heap após a última coleta de lixo for de 4MB e o GC Percentage estiver definido como 100%,
    o próximo GC será acionado quando o tamanho total do heap atingir 8MB (4MB + 100% disso, ou seja, mais 4MB, totalizando 8MB).

Quanto menor o valor do GOGC, mais vezes o GC vai ser chamado, e mais lento o programa vai ficar. Porem,
se eu aumentar o GC para 200%, o GC vai ser chamado menos vezes, e o programa vai ficar mais rápido, porém, o consumo de memoria vai ser maior.

- Quer gastar menos memoria -> Chama mais vezes o GC (Aumenta o GOGC)
- Quer gastar mais memoria -> Chama menos vezes o GC (Diminui o GOGC)

Toda vez que for chamado mais vezes o GC, sig que o programa vai parar a execução por diversas vezes, e isso pode deixar o programa mais lento.
Mas, demorar muito para chamar o GC, vai fazer com que o programa consuma mais memoria, pois o GC vai ser chamado menos vezes.

O desenvolvedor pode fazer varios testes para definir o valor do GOGC, para acompanhar a utilização do GC para entender se vale mais a pena
chamar mais vezes o GC, ou chamar menos vezes o GC.

- Comando para ver o valor do GOGC: go env GOGC
- Comando para definir o valor do GOGC: export GOGC=200
- Forma de printar o GC na tela, para ver como o GC está funcionando:
- Comando para definir e exibir rodando o GC ao rodar o programa pela linha de comando: GODEBUG=gctrace=1 GOGC=200 go run main.go
- Se eu setar o GOGC=-1 -> Desabilita o GC e ele não vai ser chamado nenhuma vez. Isso é interesante pra vc ver como o programa se comporta sem o GC.
  quanto o programa vai consumir de memoria, etc.
  - Forma de printar o GC na tela, para ver como o GC está funcionando:
    ![Captura de Tela 2024-09-24 às 07.12.43.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_HaKYUJ%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2007.12.43.png)
  - Print Alloc: Quantidade de memoria que foi alocada até o momento.
  - Print TotalAlloc: Quantidade de memoria alocada durante a execução do programa.
  - Print Sys: Quantidade de memoria alocada pelo SO.
  - Print NumGC: Quantidade de vezes que o GC foi chamado.
- É possivel debugar o GC pelo codigo/programa:
  ![Captura de Tela 2024-09-24 às 07.18.10.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_QHuPeP%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2007.18.10.png)

#### Memory Limit

![Captura de Tela 2024-09-24 às 07.20.40.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_n5TUNs%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2007.20.40.png)
![Captura de Tela 2024-09-24 às 07.28.30.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_ag1ZKY%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2007.28.30.png)

- debud.SetMemoryLimit(): É uma variavel de ambiente que é possivel definir o valor do limite de memoria do GO. Indica para o GO
  a quantidade de memoria que o programa pode utilizar.

![Captura de Tela 2024-09-24 às 07.23.15.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_YxpaM0%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2007.23.15.png)
![Captura de Tela 2024-09-24 às 07.23.47.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_TBZiOH%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2007.23.47.png)

### Channels

Channels são um mecanismo de comunicação e sincronização entre goroutines no Go. Eles permitem que goroutines troquem dados de maneira segura e eficiente, suportando a construção de programas concorrentes.

- Channels possibilita o GO trabalhar de forma concorrente e paralela.
- Channels são utilizados para fazer a comunicação e sincronização entre goroutines.
- Channels são tipos de dados, assim como int, string, etc.
- Channels são utilizados para enviar e receber dados entre goroutines.
- Channels são utilizados para passar informações entre goroutines de forma eficiente, evitando a necessidade de locks explícitos
  e reduzindo a complexidade da sincronização. (explicitos pois channels fazem locks de forma implicita, ou seja, channels fazem locks de forma automatica)

#### Dead Locks

O problema de não utilização de channel:

Problemas de sincronização

- Quando a G1 que passar um valor paga a G2, ela altera um valor na memória para a G2 ter acesso.
- O grande problema é que outras Goroutines e partes do programa também podem ter acesso aquele endereço de memória
- OU a G1 não terminou completamente de alocar o valor em memória e a G2 já fez a leitura e eventualmente uma gravação no mesmo local

Dificuldade de trabalhar com concorrência

- Data race (Race condition) -> Condição de corrida - Condição onde duas ou mais threads estão tentando acessar e manipular o mesmo valor na memoria ao mesmo tempo.
  ![Captura de Tela 2024-09-24 às 07.48.53.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_w0cTk5%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2007.48.53.png)
- Para remediar o problema utilizamos Mutex (Mutual Exclusion)
- Fazemos um lock do valor na memória e durante esse momento, somente uma Goroutine pode fazer alteraç Após isso, esse valor é liberado (unlock)
- Mutex e similares abrem muita margem para erro, pois tudo isso é feito de forma manual

Erro de Deadlocks: Se dermos lock na goroutina 1 e na goroutina 2, mas a goroutina 2 quer acessar a 1 e a 1 quer acessar a2,
vai dar deadlock, pois as duas vão ficar esperando uma a outra.

Frase que define com mais clareza a utilização de channels:

- "Do not communicate by sharing memory; instead, share memory by communicating." - Rob Pike
- "Não se comunique compartilhando memória; em vez disso, compartilhe memória se
  comunicando." - Rob Pike
- Essa frase encapsula um dos princípios fundamentais do design de sistemas concorrentes no Go. A ideia é que, ao usar channels para comunicação entre goroutines, você evita muitos dos problemas associados à concorrência e ao compartilhamento de memória direta, como condições de corrida e deadlocks.
- Em vez de várias goroutines acessarem diretamente variáveis compartilhadas (o que requer mecanismos de sincronização como locks), elas se comunicam enviando dados através de channels, o que proporciona uma maneira segura e clara de coordenar o acesso aos dados

Resumo: Toda vez que for fazer um compartilhamento direto de memoria, ou seja, goroutinas diferentes acessando a mesma variavel(endereçamento de memoria),
utilize channels para fazer a comunicação e sincronização entre as goroutinas e nao mutex, pois mutex é mais complexo de trabalhar e pode gerar deadlocks.

- Ao inves de usar mutex utilize channels, pois channels são mais faceis de trabalhar e evitam deadlocks.

#### Tipos de Channels

- Não-Bufferizados: Requerem que a operação de envio e recebimento ocorra simultaneamente. Ideal para sincronização direta entre goroutines. É possivel passar um dado/valor de cada vez. Acontece quase em tempo real o envio e recebimento de dados, ou seja, o envio e recebimento de dados ocorre ao mesmo tempo.
  ![Captura de Tela 2024-09-24 às 08.14.09.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_26Od6D%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.14.09.png)

- Bufferizados: Permitem que dados sejam armazenados temporariamente no buffer, permitindo que a goroutine de envio e a de recebimento sejam executadas em tempos diferentes. Posso eviar varios dados de uma vez para o channel, e o channel vai armazenar esses dados temporariamente no buffer, e a goroutine de recebimento vai pegar esses dados do buffer e vai processar esses dados.
  Evita que a goroutine de envio e a de recebimento fiquem bloqueadas esperando uma pela outra. Ideal para casos em que a velocidade de envio e recebimento podem variar.
  ![Captura de Tela 2024-09-24 às 08.14.52.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_FmAMOs%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.14.52.png)
  ![Captura de Tela 2024-09-24 às 08.16.46.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_ogP1YB%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.16.46.png)

#### Funcionamento do Channel

![Captura de Tela 2024-09-24 às 08.19.48.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_jOo3nd%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.19.48.png)
![Captura de Tela 2024-09-24 às 08.23.59.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_bnKRxy%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.23.59.png)
![Captura de Tela 2024-09-24 às 08.26.01.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_Lg5adG%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.26.01.png)
![Captura de Tela 2024-09-24 às 08.27.09.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_0sl1Xl%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.27.09.png)
![Captura de Tela 2024-09-24 às 08.27.53.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_aXlE1Z%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.27.53.png)

#### Quando usar Channels Bufferizados

Altas Taxas de Produção e Consumo:

- Quando há uma diferença significativa nas taxas de produção e consumo de dados entre goroutines, channels não-bufferizados podem causar gargalos.
  - Exemplo:Se uma goroutine está gerando dados muito mais rápido o que outra pode consumir, a goroutine produtora ficará frequentemente bloqueada esperando que a consumidora esteja pronta para receber, resultando em desempenho ineficiente

Pipeline de Processamento:

- Em pipelines de processamento de dados, onde os dados passam por várias etapas, cada uma implementada como uma goroutine, o uso de channels não-bufferizados pode causar bloqueios frequentes, dificultando o fluxo suave dos dados através do pipeline.
  - Exemplo: Se uma goroutine está gerando dados muito mais rápido do que outra pode consumir, a goroutine produtora ficará frequentemente bloqueada esperando que a consumidora esteja pronta para receber, resultando em desempenho ineficiente

Tarefas Assíncronas:

- Quando se trabalha com tarefas que não precisam ser sincronizadas estritamente, como registros de logs ou coleta de métricas, usar channels não-bufferizados pode introduzir latência desnecessária.
  - Exemplo: Ao registrar logs em um servidor de alta carga, esperar que cada log seja processado antes de prosseguir pode impactar negativamente a performance. Um channel bufferizado permite que a produção de logs continue sem esperar pelo processamento de cada mensagem de log.

Comunicação Entre Múltiplas Goroutines:

- Quando há comunicação entre várias goroutines produtoras e consumidoras, channels não- bufferizados podem aumentar a contenção e reduzir a paralelização efetiva.
  - bExemplo: Se várias goroutines tentam enviar dados para um único channel não-bufferizado, elas competirão pelo acesso, resultando em bloqueios frequentes e menor desempenho.

Operações de I/O(Input/Output)

- Em operações de I/O, onde o tempo de espera pode ser significativo, o uso de channels não- bufferizados pode levar a bloqueios desnecessários.
  - Exemplo: Uma goroutine que lê dados de um arquivo e outra que escreve esses dados em uma rede. Se a escrita na rede for mais lenta, a leitura do arquivo será bloqueada frequentemente

Resumo dos Problemas com Channels Não-Bufferizados:

- Bloqueios Frequentes: Podem resultar em bloqueios frequentes entre produtor e consumidor. • Gargalos de Desempenho: Ineficiência em casos de alta produção ou processamento lento.
- Latência: Introduz latência desnecessária em tarefas assíncronas.
- Contenção: Aumenta a contenção entre múltiplas goroutines.

#### Quando usar Channels Não-Bufferizados

Sincronização Estrita:

- Quando é necessário garantir que a produção e o consumo de dados ocorrem de maneira sincronizada, channels não-bufferizados são ideais. Eles garantem que cada operação de envio é emparelhada diretamente com uma operação de recebimento.
  ![Captura de Tela 2024-09-24 às 08.53.28.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_nw7DaK%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.53.28.png)

Handshake (Aperto de Mão):

- Em situações onde duas goroutines precisam trocar informações ou confirmar que uma ação foi realizada antes de prosseguir, channels não-bufferizados fornecem um mecanismo simples e seguro.
  ![Captura de Tela 2024-09-24 às 08.53.53.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_79aZfi%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.53.53.png)

Eventos Temporizados:

- Quando você precisa lidar com eventos temporizados, como timeouts, channels não-bufferizados podem ser usados em conjunto com o select para implementar timeouts de maneira simples e eficaz.
  ![Captura de Tela 2024-09-24 às 08.54.23.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_NISbBF%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.54.23.png)

Coordenar Finalização:

- Quando várias goroutines precisam ser coordenadas para garantir que todas elas completem antes que o programa possa prosseguir, channels não-bufferizados podem ser utilizados para sinalizar a conclusão.
  ![Captura de Tela 2024-09-24 às 08.54.44.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Ftw%2F24n8jlwx0195szqp43qk52xjbh8kn2%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_x09PsV%2FCaptura%20de%20Tela%202024-09-24%20%C3%A0s%2008.54.44.png)

#### Qual o número mágico para buffers de channels?

Parametrização do tamanho do Buffer

- Taxa de produção e consumo

  - Taxa Variável: Se a taxa de produção e consumo varia significativamente, um buffer maior pode ajudar a suavizar as diferenças e evitar bloqueios frequentes.
  - Taxa Constante: Se as taxas de produção e consumo são constantes e iguais, um buffer menor pode ser suficiente.

- Latência e Desempenho

  - Baixa Latência: Se a latência é crítica, um buffer menor pode ser preferível para garantir que os dados
    sejam processados rapidamente.
  - Alto Desempenho: Um buffer maior pode ajudar a aumentar o desempenho em sistemas de alta taxa de transferência, reduzindo a contenção entre goroutines.

- Memória Disponível

  - Uso de Memória: Buffers grandes consomem mais memória. Certifique-se de que há memória suficiente disponível e considere o impacto no uso geral do sistema.
  - Padronização de pipelines
  - Etapas de Pipeline: Se você está construindo um pipeline de processamento, cada etapa pode se beneficiar de um buffer que permite processar dados em lotes, melhorando a eficiência geral

- Resumo: Não há um tamanho de buffer único que seja ideal para todos os casos. A escolha do tamanho do buffer deve ser baseada em:
  - Taxa de produção e consumo.
  - Requisitos de latência e desempenho.
  - Memória disponível.
  - Características específicas do seu pipeline de processamento.
  - Monitorar e ajustar o tamanho do buffer conforme necessário, juntamente com testes de desempenho,
  - pode ajudar a encontrar o equilíbrio certo para seu caso de uso específico.
