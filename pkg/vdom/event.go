package vdom

// EventType is the DOM event type
type EventType string

// HTML/SVG event handler types
const (
	Click       EventType = "click"
	MouseDown             = "mousedown"
	MouseUp               = "mouseup"
	MouseEnter            = "mouseenter"
	MouseLeave            = "mouseleave"
	MouseMove             = "mousemove"
	ContextMenu           = "contextmenu"
	TouchStart            = "touchstart"
	TouchEnd              = "touchend"
	KeyDown               = "keydown"
	KeyUp                 = "keyup"
	Change                = "change"
)

// HandlerFunc is the registered callback method
type HandlerFunc func(*Element, *Event)

// EventHandler is the DOM registered event handler
type EventHandler struct {
	Type        string
	handlerFunc HandlerFunc
}

// MakeEventHandler creates a EventHanlder with a listener function
func MakeEventHandler(eventType EventType, handlerFunc HandlerFunc) EventHandler {
	return EventHandler{Type: string(eventType), handlerFunc: handlerFunc}
}

// Event is the event data
type Event struct {
	Type string
	Data map[string]interface{}
}
