package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int    `json:"numero"`
	Saldo  int    `json:"s" validate:"gt=0"` //temos tags de validação tbm
	Senha  string `json:"-"`                 //"-" omit a informação
}

func main() {
	//processo de serilizacao e deserilizacao de jsons
	conta := Conta{Numero: 1, Saldo: 100}
	//transforma oque está na struct em json
	// Forma 1 - guardando em uma variavel
	res, err := json.Marshal(conta) //Serializar
	if err != nil {
		println(err)
	}
	println(string(res)) //json sempre vai retornar em bytes, por isso convertemos para string

	// Forma 2 -retornando direto (encoder)
	//NewEncoder: ja faz o processo de serilização entregando pra alguém
	//encoder: é um cara que recebe o valor e faz um encoding gravando em outro lugar. Pode ser um arquivo, tela do terminal ...
	// encoder := json.NewEncoder(os.Stdout) //os.Stdout: sai com a saida padrao no terminal
	// encoder.Encode(conta) //forma 1

	// encoder := json.NewEncoder(os.Stdout).Encode(conta) //forma 2

	// json.NewEncoder(os.Stdout).Encode(conta) //forma 3

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	//decoded: processso ao contrario
	//transforma json em struct
	//forma sem uso de tag json na struct
	// jsonPuro := []byte(`{"Numero":2, "Saldo":200}`) //por padrao o json é sempre no formato de slice de bites
	// var contaX Conta
	// err = json.Unmarshal(jsonPuro, &contaX)
	// if err != nil {
	// 	println(err)
	// }

	// println(contaX.Saldo)

	//uso do retorno do json referente as tags da struct
	//forma com uso de tag json na struct
	jsonPuro2 := []byte(`{"numero":2, "s":200}`) //por padrao o json é sempre no formato de slice de bites
	var contaX2 Conta
	err = json.Unmarshal(jsonPuro2, &contaX2)
	if err != nil {
		println(err)
	}

	fmt.Println(contaX2)

	fmt.Println(contaX2.Numero)
	fmt.Println(contaX2.Saldo)

}
