package vdom

// EventType is the DOM event type
type EventType int

// XML event handler types
const (
	Click EventType = iota
	MouseDown
	MouseUp
	KeyDown
	KeyUp
)

// Listener is the registered callback method
type Listener func(*Element, *Event)

// EventHandler is the DOM registered event handler
type EventHandler struct {
	Type     EventType
	Listener Listener
}

// NewEventHandler creates an EventHanlder with a listener function
func NewEventHandler(eventType EventType, listener Listener) EventHandler {
	return EventHandler{Type: eventType, Listener: listener}
}

// Event is the event data
type Event struct {
	Type EventType
	Rune rune
}
