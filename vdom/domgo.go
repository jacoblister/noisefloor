//+build !js

package vdom

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//applyPatchToDom applies the patch for the GoLang native target
func applyPatchToDom(patch *Patch) {
	// fmt.Println("GoLang target apply patch", patch)
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

type msg struct {
	Num int
}

func clientProcess(conn *websocket.Conn) {
	patch := fullDomPatch()
	conn.WriteJSON(patch)

	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
			return
		}
	}
}

//ListenAndServe begins and HTTP server for the application
func ListenAndServe() {
	// patch := fullDomPatch()
	// r, _ := json.Marshal(patch)
	// fmt.Println(string(r))

	println("done")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/client", clientHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
