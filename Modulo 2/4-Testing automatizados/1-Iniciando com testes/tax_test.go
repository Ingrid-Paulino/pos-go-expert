package tax

import "testing"

// Comando para rodar o teste: go test . OU go test -v
// Vai passar
func TestCalculateTax(t *testing.T) {
	value := 500.0
	expected := 5.0

	result := CalculateTax(value)

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

// Vai falhar
func TestCalculateTaxError(t *testing.T) {
	value := 500.0
	expected := 6.0

	result := CalculateTax(value)

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}
