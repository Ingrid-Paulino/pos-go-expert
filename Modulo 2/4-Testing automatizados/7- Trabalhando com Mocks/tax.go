package tax

import "errors"

func CalculateTax(value float64) (float64, error) {
	if value <= 0 {
		return 0, errors.New("value must be greater than 0")
	}

	if value >= 1000 && value < 20000 {
		return 10.0, nil
	}

	if value >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}

/*
	Vamos imaginar que estamos em um momento da aplicação que vamos calcular a taxa e salvar o valor do calculo em um banco de dados

mas não tenho ainda um banco de dados e não sei qual banco vou utilizar. No entanto, quero testar essa minha função que vai calcular e salvar no banco de dados.
Com o Mock eu posso simular o comportamento do banco de dados e testar a minha função.
*/

type Repository interface {
	SaveTax(tax float64) error
}

// O objetivo dessa função não é testar o SaveTax, pois ele será testado em outra camada. A intenção é testar o CalculateAndSaveTax simulando/mockando o comportamento do SaveTax.
func CalculateAndSaveTax(value float64, repository Repository) error { //repository é so uma interface que tem o metodo SaveTax, não temos ainda uma struct que implementa essa interface e que se conecta no banco de dados
	tax := CalculateTax2(value)

	// save in database
	return repository.SaveTax(tax)
}

func CalculateTax2(value float64) float64 {
	if value <= 0 {
		return 0
	}

	if value >= 1000 && value < 20000 {
		return 10.0
	}

	if value >= 20000 {
		return 20.0
	}
	return 5.0
}
