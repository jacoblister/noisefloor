//+build !js

package vdom

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type eventHandlerKey struct {
	id        string
	eventType EventType
}

type eventHandlerValue struct {
	element      *Element
	eventHandler *EventHandler
}

type domEvent struct {
	Type      EventType
	ElementID string
	Data      string
}

var eventHandlerMap map[eventHandlerKey]eventHandlerValue

//updateEventHandlersRecursive set all event handlers in the element tree
func updateEventHandlersRecursive(element *Element) {
	for _, handler := range element.EventHandlers {
		id := element.Attrs["id"].(string)
		eventHandlerKey := eventHandlerKey{id: id, eventType: handler.Type}
		eventHandlerValue := eventHandlerValue{element: element, eventHandler: &handler}
		eventHandlerMap[eventHandlerKey] = eventHandlerValue
	}

	for _, child := range element.Children {
		updateEventHandlersRecursive(&child)
	}
}

var activeConnections map[*websocket.Conn]int

//applyPatchToDom applies the patch for the GoLang native target
func applyPatchToDom(patch *Patch) {
	eventHandlerMap = map[eventHandlerKey]eventHandlerValue{}
	updateEventHandlersRecursive(&patch.Element)
}

//rootHandler servers to main dom transfer container and script
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../domgo.html")
}

//handleDomEvent processes a DOM event received through a connection
func handleDomEvent(domEvent domEvent) {
	eventHandlerKey := eventHandlerKey{id: domEvent.ElementID, eventType: domEvent.Type}

	handler := eventHandlerMap[eventHandlerKey]
	event := Event{Type: domEvent.Type, Data: domEvent.Data}
	handler.eventHandler.handlerFunc(handler.element, &event)

	for conn := range activeConnections {
		println("connection event:", conn)
		patch := fullDomPatch()
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
		handleDomEvent(msg)
	}
}

//ListenAndServe begins and HTTP server for the application
func ListenAndServe() {
	activeConnections = map[*websocket.Conn]int{}
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/client", clientHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
