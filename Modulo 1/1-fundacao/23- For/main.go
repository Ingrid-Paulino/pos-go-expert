package main

func main() {
	//temos apenas o for em GO
	//forma 1
	for i := 0; i < 10; i++ {
		println(i)
	}

	//forma 2
	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}

	//forma 3
	i := 0
	for i < 5 {
		println(i)
		i++
	}

	//forma 4 -> loop infinito
	for {
		println("Hello, World!")
	}

	//funciona na versão 1.22.1
	for x := range 10 {
		println("Hello, World!")
	}
}
