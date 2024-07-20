package tax

import (
	"testing"

	"github.com/stretchr/testify/assert" //biblioteca para testes. Facilita a escrita de testes
)

// Comando para rodar o teste: go test . OU go test -v
// Comando para rodar o teste com cobertura de código - coverage:
//- go test -coverprofile=coverage.out --> gera o arquivo coverage.out
//- go test -coverprofile=coverage.out && go tool cover -html=coverage.out --> consigo ver o relatório em html. Me mostra linhas cobertas e linhas não cobertas

func TestCalculateTaxBatch(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Equal(t, 10.0, tax)
	assert.Nil(t, err)

	tax, err = CalculateTax(20000.0)
	assert.Equal(t, 20.0, tax)
	assert.Nil(t, err)

	tax, err = CalculateTax(20)
	assert.Equal(t, 5.0, tax)
	assert.Nil(t, err)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "value must be greater than 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}
