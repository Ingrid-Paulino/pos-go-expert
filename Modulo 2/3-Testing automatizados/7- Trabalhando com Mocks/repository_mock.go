package tax

import "github.com/stretchr/testify/mock"

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(value float64) error {
	args := m.Called(value)
	return args.Error(0)
	//OBS: se eu retornasse nessa funcao um int, string, float ou qualquer outro tipo, eu teria que usar args.Get(0).(tipo) para fazer o cast ou
	//return args.Int(0)
	//return args.String(0)
	//return args.Float(0)
}
