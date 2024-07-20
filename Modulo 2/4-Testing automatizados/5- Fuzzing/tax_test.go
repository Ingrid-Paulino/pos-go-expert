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
		result := calculateTax(test.value)
		if result != test.expected {
			t.Errorf("Expected %f, but got %f", test.expected, result)
		}
	}
}

// Fuzzing trabalha com valores aleatórios para testar a função. A ideia é que o Fuzzing encontre erros que não seriam encontrados com testes unitários normais.
// O fazzing gera valores aleatorios para testar a funcao, assim testa uma gama maior de valores, posiblitando encontrar valores que quebrem a funcao.
// Fazzing testa minha aplicação com varios valores diferentes, ate tentar quebrar minha aplicação.
// OBS: fazzing cria a pasta testdata automaticamente, com o valor que ele testou e não passou no teste. Facilitando pegar erros de logicas que nao foram tratados. Depois que arrumar a logica o fazzing indica um comando ex: go test -run=FuzzCalculateTax/e953ca6392db88ba para rodar o arquivo geradp por ele e testar o valor que não estava contemplado na logica

// para rodar o Fuzzing:
// go test -fuzz=.
// go test -fuzz=. -fuzztime=10s
// go test -fuzz=FuzzCalculateTax -fuzztime=10s -run=ˆ=
// go test -fuzz=. -run=ˆ# -> roda so o fuzzing e nao os outros testes pois nenhum tem o prefixo # no nome
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0} //valores que quero testar na funcao
	for _, value := range seed {
		f.Add(value) //adiciona o valor ao fuzzing para testar. Essa funcao pode receber n parametros
	}

	f.Fuzz(func(t *testing.T, value float64) { //funcao que vai rodar o fuzzing, Essa funcao pode receber n parametros e o primeiro sempre é o *testing.T e os demais sao os parametros que adicionei no Add
		result := calculateTax(value)
		if value > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 0", result)
		}

		if value <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})
}
