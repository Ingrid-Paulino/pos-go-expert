package main

import (
	"fmt"

	//eventsPrivado "github.com/Ingrid-Paulino/fcutils/pkg/events" //repositorio privado: não é posivel importar o pacote como fazemos normalmente com os pacotes publicos
	"github.com/devfullcycle/fcutils/pkg/events" //repositorio publico: consigo importar o pacote sem problemas
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
	//ed2 := eventsPrivado.NewEventDispatcher()
	//fmt.Println(ed2)

}
