package tax

import "testing"

// Comando para rodar o teste: go test . OU go test -v
// Comando para rodar o teste com cobertura de código - coverage:
//- go test -coverprofile=coverage.out --> gera o arquivo coverage.out
//- go test -coverprofile=coverage.out && go tool cover -html=coverage.out --> consigo ver o relatório em html. Me mostra linhas cobertas e linhas não cobertas

// Posso rodar separados tbm:
//- go test -coverprofile=coverage.out --> gera o arquivo coverage.out
//- go tool cover -html=coverage.out

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		value, expected float64
	}

	//Testa 3 informacoes diferentes ao mesmo tempo
	table := []calcTax{
		{value: 500.0, expected: 5.0},
		{value: 1000.0, expected: 10.0},
		{value: 1500.0, expected: 10.0},
	}

	for _, test := range table {
		result := CalculateTax(test.value)
		if result != test.expected {
			t.Errorf("Expected %f, but got %f", test.expected, result)
		}
	}
}

/*
Benchmark é um teste de desempenho que mede o tempo de execução de uma função.
Para criar um benchmark, você precisa adicionar um prefixo "Benchmark" ao nome da função e passar um parâmetro *testing.B.
O parâmetro *testing.B é usado para controlar o número de iterações do benchmark e relatar os resultados.
Você pode usar o método b.N para obter o número de iterações do benchmark.

banchmark é para verificar/medir performance de uma função.

banchmark é executado com o comando: go test -bench=.

No exemplo abaixo, criamos um benchmark para a função calculateTax.
*/

func BenchmarkCalculateTax(b *testing.B) { //operacao mais rapida
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// Esse teste vai demorar 1 segundo para rodar cada iteração do benchmark.
func BenchmarkCalculateTax2(b *testing.B) { //operação mais lenta
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

//banchmark é executado com o comando: go test -bench=.
// para rodar apenas o benchmark sem os testes: go test -bench=. -run=ˆ#
// Vai rodar 10 vezes cada branchMarck: go test -bench=. -run=ˆ# -count=10
// Vai rodar em 3 seg o branchMarck: go test -bench=. -run=ˆ# -count=10 -benchtime=3s --> aqui vai atrasar pq a segunda função tem um sleep de 1s a cada execução
// Comando mostra quantidade de alocação de memoria que a função esta gastando: go test -bench=. -run=ˆ# -benchmem

/*
Retorno:
goos: darwin -> sistema operacional mac
goarch: arm64 -> arquitetura do processador
pkg: taxgo -> pacote sendo testado
BenchmarkCalculateTax-8  -> Nome da função testada e o traço 8 é a quantidade de nucleos computacionais que foi utilizado para rodar essa função. Que basicamente é o número de CPUs disponíveis para o benchmark na minha maquina.
1000000000  -> 1 bilhão de iterações / Quantidade de iterações que foi executada na função
0.3682 ns/op -> Tempo médio de execução por iteração. Nesse caso, 0.3682 nanosegundos por iteração.
*/
