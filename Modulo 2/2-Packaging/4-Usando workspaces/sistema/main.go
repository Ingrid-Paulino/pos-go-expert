package main

/*
Obs: Diferente da aula anterior, agora o import não foi feito com o sistema do replace para rodar local e sim com workspace. Detalhes no arquivo note.txt
*/
import (
	"github.com/google/uuid"
	"github.com/ingrid-paulino/goexpert/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add2())
	println(uuid.New().String())
}
