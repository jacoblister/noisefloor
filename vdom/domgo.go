//+build !js

package vdom

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func updateEventHandlersRecursive(element *Element) {
	for _, child := range element.Children {
		updateEventHandlersRecursive(&child)
	}
}

//applyPatchToDom applies the patch for the GoLang native target
func applyPatchToDom(patch *Patch) {
	updateEventHandlersRecursive(&patch.Element)
	fmt.Println("GoLang target apply patch", patch)
}

//rootHandler servers to main dom transfer container and script
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../domgo.html")
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

	go clientProcess(conn)
}

type domEvent struct {
 	ElementID string
	Data      string
}

func clientProcess(conn *websocket.Conn) {
	patch := fullDomPatch()
	conn.WriteJSON(patch)

	for {
		msg := domEvent{}

		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", msg)
	}
}

//ListenAndServe begins and HTTP server for the application
func ListenAndServe() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/client", clientHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
