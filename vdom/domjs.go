//+build js

package vdom

import "github.com/gopherjs/gopherjs/js"

func addEventHandler(element *Element, domNode *js.Object, handler *EventHandler) {
	domNode.Call("addEventListener", "click", func(jsEvent *js.Object) {
		event := Event{Type: Click}
		handler.Listener(element, &event)
	})
}

func createElementRecursive(element *Element, domNode *js.Object) {
	document := js.Global.Get("document")

	switch element.Type {
	case Normal:
		child := document.Call("createElement", element.Name)

		for _, attr := range element.Attrs {
			child.Call("setAttribute", attr.Name, attr.Value)
		}

		for _, handler := range element.eventHandlers {
			addEventHandler(element, domNode, &handler)
		}

		for _, elem := range element.children {
			createElementRecursive(&elem, child)
		}

		domNode.Call("appendChild", child)
	case Text:
		domNode.Set("innerText", element.Attrs[0].Value)
	}
}

func applyPatchToDom(patch *Patch) {
	root := js.Global.Get("document").Get("body")
	createElementRecursive(&patch.element, root)
}
