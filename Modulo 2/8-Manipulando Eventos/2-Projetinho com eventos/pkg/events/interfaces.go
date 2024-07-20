package events

import (
	"sync"
	"time"
)

// Evento
type EventInterface interface {
	GetName() string         // Get the name of the event
	GetDataTime() time.Time  // Get the date and time of the event
	GetPayload() interface{} // Get the payload of the event, por ser uma interface vazia, qualquer coisa pode implementar essa funcao. Payload: são os dados que tem no evento
}

// Operaçoes que serão executadas quando um evento é chamado
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup) //handle method é a unica funcao obrigatoria, pois ele que executa as operacoes
}

// Gerenciador dos nossos eventos/operaçoes
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error // Register an event and its handler
	Dispatch(event EventInterface) error                            // Dispatch an event to its handlers
	Remove(eventName string, handler EventHandlerInterface) error   // Remove an event and its handler
	Has(eventName string, handler EventHandlerInterface) bool       // Check if an event has a handler
	Clear() error                                                   // Limpa o event dispatcher matando todos os eventos que estao registrados nele
}
