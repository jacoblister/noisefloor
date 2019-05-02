//+build js

package vdom

import (
	"github.com/gopherjs/gopherjs/js"
)

func addEventHandler(element *Element, domNode *js.Object, handler *EventHandler) {
	domNode.Call("addEventListener", handler.Type, func(jsEvent *js.Object) {
		var eventData string
		switch handler.Type {
		case "change":
			eventData = jsEvent.Get("target").Get("value").String()
		case "click",
			"mousedown",
			"mouseup",
			"mouseenter",
			"mouseleave":
			eventData = jsEvent.Get("buttons").String()
		}

		event := Event{Type: handler.Type, Data: eventData}

		updateDomBegin()
		handler.handlerFunc(element, &event)
		patch := updateDomEnd()
		applyPatchToDom(patch)
	})
}

func createElementRecursive(svgNamespace bool, element *Element) *js.Object {
	document := js.Global.Get("document")

	var node *js.Object
	if svgNamespace == true {
		node = document.Call("createElementNS", "http://www.w3.org/2000/svg", element.Name)
	} else {
		node = document.Call("createElement", element.Name)
	}

	for name, value := range element.Attrs {
		node.Call("setAttribute", name, value)
	}

	for _, handler := range element.EventHandlers {
		addEventHandler(element, node, &handler)
	}

	for _, child := range element.Children {
		switch child.Type {
		case Normal:
			childNode := createElementRecursive(svgNamespace, &child)
			node.Call("appendChild", childNode)
		case Text:
			node.Set("innerText", child.Attrs["Text"])
		}
	}

	return node
}

func applyPatchToDom(patch *Patch) {
	switch patch.Type {
	case Replace:
		root := js.Global.Get("document").Get("body")

		child := root.Get("lastElementChild")
		if child != nil {
			root.Call("removeChild", child)
		}

		node := createElementRecursive(patch.SVGNamespace, &patch.Element)
		root.Call("appendChild", node)
	}
}

//componentUpdateListen reads and applies background component state changes
func componentUpdateListen(c chan Component) {
	for {
		component := <-c

		updateDomBegin()
		UpdateComponent(component)
		patch := updateDomEnd()
		applyPatchToDom(patch)
	}
}

//ListenAndServe starts the javascript target
func ListenAndServe() {
	componentUpdate = make(chan Component, 10)
	go componentUpdateListen(componentUpdate)

	applyPatchToDom(fullDomPatch())
}
