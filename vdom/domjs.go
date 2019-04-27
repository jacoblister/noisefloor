//+build js

package vdom

import "github.com/gopherjs/gopherjs/js"

func addEventHandler(element *Element, domNode *js.Object, handler *EventHandler) {
	switch handler.Type {
	case MouseDown:
		domNode.Call("addEventListener", "mousedown", func(jsEvent *js.Object) {
			event := Event{Type: MouseDown}
			handler.handlerFunc(element, &event)
		})
	case MouseUp:
		domNode.Call("addEventListener", "mouseup", func(jsEvent *js.Object) {
			event := Event{Type: MouseUp}
			handler.handlerFunc(element, &event)
		})
	case Click:
		domNode.Call("addEventListener", "click", func(jsEvent *js.Object) {
			event := Event{Type: Click}
			handler.handlerFunc(element, &event)
		})
	case Change:
		domNode.Call("addEventListener", "change", func(jsEvent *js.Object) {
			value := jsEvent.Get("target").Get("value").String()
			event := Event{Type: Change, Data: value}
			handler.handlerFunc(element, &event)
		})
	}
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

//ListenAndServe is a no-op for the javascript target
func ListenAndServe() {}
