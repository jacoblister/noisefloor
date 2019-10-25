//+build !js

package vdom

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jacoblister/noisefloor/pkg/vdom/assets"
)

type eventHandlerKey struct {
	id        string
	eventType string
}

type eventHandlerValue struct {
	element      *Element
	eventHandler *EventHandler
	Type         string
}

type domEvent struct {
	Type      string
	ElementID string
	Data      map[string]interface{}
}

var domUpdateMutex sync.Mutex
var eventHandlerMap map[eventHandlerKey]eventHandlerValue

//updateEventHandlersRecursive set all event handlers in the element tree
func updateEventHandlersRecursive(element *Element) {
	for i := 0; i < len(element.EventHandlers); i++ {
		handler := &element.EventHandlers[i]
		id, ok := element.Attrs["id"].(string)
		if !ok {
			panic("Must define id for element with event handler")
		}
		key := eventHandlerKey{id: id, eventType: handler.Type}
		value := eventHandlerValue{element: element, eventHandler: handler, Type: handler.Type}

		_, present := eventHandlerMap[key]
		if present {
			panic("Duplicate element key for event handler: " + id)
		}
		eventHandlerMap[key] = value
	}

	for i := 0; i < len(element.Children); i++ {
		updateEventHandlersRecursive(&element.Children[i])
	}
}

var activeConnections map[*websocket.Conn]int

//applyPatchToDom applies the patch for the GoLang native target
func applyPatchToDom(patchList PatchList) {
	for i := 0; i < len(patchList.Patch); i++ {
		eventHandlerMap = map[eventHandlerKey]eventHandlerValue{}
		updateEventHandlersRecursive(&patchList.Patch[i].Element)
	}
}

// domEventDataTranslate converts event data types (only use integers at the moment)
func domEventDataTranslate(eventData map[string]interface{}) map[string]interface{} {
	translatedData := map[string]interface{}{}
	for key, value := range eventData {
		switch value.(type) {
		case float32:
			translatedData[key] = int(value.(float32))
		case float64:
			translatedData[key] = int(value.(float64))
		default:
			translatedData[key] = value
		}
	}
	return translatedData
}

//handleDomEvent processes a DOM event received through a connection
func handleDomEvent(domEvent domEvent) {
	eventHandlerKey := eventHandlerKey{id: domEvent.ElementID, eventType: domEvent.Type}

	handler := eventHandlerMap[eventHandlerKey]
	event := Event{Type: domEvent.Type, Data: domEventDataTranslate(domEvent.Data)}

	updateDomBegin()
	handler.eventHandler.handlerFunc(handler.element, &event)
	patch := updateDomEnd()
	applyPatchToDom(fullDomPatch()) // TODO - improve this, not efficient, should apply new patch, not full patch

	for conn := range activeConnections {
		conn.WriteJSON(patch)
	}
}

//clientHandler allows a websocket client to connect
func clientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	activeConnections[conn] = 1

	go clientProcess(conn)
}

func clientProcess(conn *websocket.Conn) {
	defer func() {
		println("Closed Connection")
		delete(activeConnections, conn)
		conn.Close()
	}()

	patch := fullDomPatch()
	conn.WriteJSON(patch)

	for {
		msg := domEvent{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			println("conn read error")
			break
		}
		domUpdateMutex.Lock()
		handleDomEvent(msg)
		domUpdateMutex.Unlock()
	}
}

//componentUpdateListen reads and applies background component state changes
func componentUpdateListen(c chan Component) {
	for {
		component := <-c

		domUpdateMutex.Lock()
		updateDomBegin()
		UpdateComponent(component)
		patch := updateDomEnd()
		applyPatchToDom(fullDomPatch()) // TODO - improve this, not efficient, should apply new patch, not full patch
		domUpdateMutex.Unlock()

		for conn := range activeConnections {
			conn.WriteJSON(patch)
		}
	}
}

//ListenAndServe begins and HTTP server for the application
func ListenAndServe(args ...interface{}) {
	applyPatchToDom(fullDomPatch())

	activeConnections = map[*websocket.Conn]int{}

	componentUpdate = make(chan Component, 10)
	go componentUpdateListen(componentUpdate)

	// fs := http.FileServer(http.Dir("../../assets/files"))
	fs := http.FileServer(assets.Assets)
	http.Handle("/", fs)
	http.HandleFunc("/client", clientHandler)

	// Add optional resource handlers
	if len(args) == 2 {
		prefix := args[0].(string)
		resourceFS := args[1].(http.FileSystem)

		http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(resourceFS)))
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
