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
	Change
)

// HandlerFunc is the registered callback method
type HandlerFunc func(*Element, *Event)

// EventHandler is the DOM registered event handler
type EventHandler struct {
	Type        EventType
	Component   Component
	HandlerFunc HandlerFunc
}

// MakeEventHandler creates a EventHanlder with a listener function
func MakeEventHandler(eventType EventType, handlerFunc HandlerFunc) EventHandler {
	return EventHandler{Type: eventType, HandlerFunc: handlerFunc}
}

// MakeComponentEventHandler creates a EventHanlder with a component receiver
func MakeComponentEventHandler(eventType EventType, component Component, handlerFunc HandlerFunc) EventHandler {
	return EventHandler{Type: eventType, Component: component, HandlerFunc: handlerFunc}
}

// Event is the event data
type Event struct {
	Type EventType
	Data string
}
