package tax

import (
	"errors"
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

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	//se receber a taxa 0 vamos querer que retorne um erro e não salve no banco de dados
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	//repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax")) //mock.Anything é um coringa. Vai aceitar qualquer coisa
	//repository.On("SaveTax", 10.0).Return(nil).Once()                              //Once é para dizer que essa chamada só pode ser feita uma vez

	err := CalculateAndSaveTax(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateAndSaveTax(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)            //verifica se todas as expectativas foram atendidas
	repository.AssertCalled(t, "SaveTax", 10.0) //verifica se o método foi chamado com os argumentos corretos
	//repository.AssertNumberOfCalls(t, "SaveTax", 3) //verifica o número de chamadas para um método específico
}
