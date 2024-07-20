package tax

import "time"

func CalculateTax(value float64) float64 {
	if value == 0 {
		return 0
	}
	if value >= 1000 {
		return 10.0
	}
	return 5.0
}

func CalculateTax2(value float64) float64 {
	time.Sleep(time.Millisecond) // Simula um atraso na execução da função
	if value == 0 {
		return 0
	}
	if value >= 1000 {
		return 10.0
	}
	return 5.0
}
