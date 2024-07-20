package tax

import "testing"

// Comando para rodar o teste: go test . OU go test -v
// O Batch testa varios valores ao mesmo tempo
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
