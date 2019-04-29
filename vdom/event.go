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

func (e EventType) String() string {
	return [...]string{
		"click",
		"mousedown",
		"mouseup",
		"keydown",
		"keyup",
		"change",
	}[e]
}

// HandlerFunc is the registered callback method
type HandlerFunc func(*Element, *Event)

// EventHandler is the DOM registered event handler
type EventHandler struct {
	Type        string
	handlerFunc HandlerFunc
}

// MakeEventHandler creates a EventHanlder with a listener function
func MakeEventHandler(eventType EventType, handlerFunc HandlerFunc) EventHandler {
	return EventHandler{Type: eventType.String(), handlerFunc: handlerFunc}
}

// Event is the event data
type Event struct {
	Type string
	Data string
}
