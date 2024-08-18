package events

import (
	"errors"
	"sync"
)

//OBS: A pasta events tem uma forma de criar o sistema de eventos, mas é uma forma mais complexa. Como o kafka e o rabbitmq ja tem uma forma de criar o sistema de eventos pronto.

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Dispatch dispara um evento para todos os handlers registrados - Faz todos os handlers atrelados a um evento serem executados
func (ev *EventDispatcher) Dispatch(event EventInterface) error {
	//Pecorre o map de handlers e executa cada um deles
	if handlers, ok := ev.handlers[event.GetName()]; ok { //verifica se o evento ja foi registrado
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			// Usando a goroutine para executar cada handler em paralelo, todos rodaram ao mesmo tempo e ganharemos performance
			// é importante usarmos o WaitGroup para garantir que todas as goroutines terminem antes de retornar
			// se não colacar o waitgroup os testes vão falhar, pois o teste vai terminar antes de todas as goroutines terminarem
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

// Registra um evento e seu handler
func (d *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := d.handlers[eventName]; ok { //verifica se o evento ja foi registrado
		for _, h := range d.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	d.handlers[eventName] = append(d.handlers[eventName], handler)
	return nil
}

// Verifica se evento existe no map de handlers
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok { //verifica se o evento ja foi registrado
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok { //verifica se o evento ja foi registrado
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				//O primeiro valor do apend é a posisao 0, o segundo é o valor que eu quero adicionar - explicação disso no Modulo 1/1-fundacao/6-Slices/ex2
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...) //remove o handler do slice
				return nil
			}
		}
	}
	return nil
}

// Limpa todos os eventos do map handlers
func (d *EventDispatcher) Clear() {
	d.handlers = make(map[string][]EventHandlerInterface)
}
