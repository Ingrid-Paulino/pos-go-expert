package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Registra um evento e seu handler
func (d *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := d.handlers[eventName]; ok {
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
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

// Limpa todos os eventos do map handlers
func (d *EventDispatcher) Clear() {
	d.handlers = make(map[string][]EventHandlerInterface)
}
