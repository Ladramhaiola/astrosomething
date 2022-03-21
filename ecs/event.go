package ecs

type EventType uint

type Event interface {
	Type() EventType
}

type EventManager struct {
	listeners map[EventType][]func(Event)
}

// TODO: remove listener & some more complicated ES :)
func (em *EventManager) AddListener(t EventType, listeners ...func(Event)) {
	em.listeners[t] = append(em.listeners[t], listeners...)
}

func (em *EventManager) SendEvent(event Event) {
	for _, listener := range em.listeners[event.Type()] {
		listener(event)
	}
}

func NewEventManager() *EventManager {
	return &EventManager{listeners: make(map[EventType][]func(Event))}
}
