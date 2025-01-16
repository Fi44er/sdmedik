package events

import "sync"

type EventType string

const (
	EventDataCreatedOrUpdated EventType = "data_created_or_updated"
	EventDataDeleted          EventType = "data_delete"
)

// Event представляет собой событие
type Event struct {
	Type     EventType
	Data     interface{}
	DataType string
}

// EventBus управляет подписками и рассылкой событий
type EventBus struct {
	subscribers map[EventType][]func(Event)
	lock        sync.RWMutex
}

// NewEventBus создает новый экземпляр EventBus
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventType][]func(Event)),
	}
}

// Subscribe подписывает обработчик на определенный тип события
func (eb *EventBus) Subscribe(eventType EventType, handler func(Event)) {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	eb.subscribers[eventType] = append(eb.subscribers[eventType], handler)
}

// Publish публикует событие
func (eb *EventBus) Publish(event Event) {
	eb.lock.RLock()
	defer eb.lock.RUnlock()

	for _, handler := range eb.subscribers[event.Type] {
		go handler(event) // Обработчики выполняются асинхронно
	}
}
