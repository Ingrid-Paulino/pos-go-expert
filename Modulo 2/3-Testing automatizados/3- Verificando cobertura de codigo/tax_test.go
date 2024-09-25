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
