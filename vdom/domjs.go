//+build js

package vdom

import "github.com/gopherjs/gopherjs/js"

func addEventHandler(element *Element, domNode *js.Object, handler *EventHandler) {
	switch handler.Type {
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

func createElementRecursive(element *Element, domNode *js.Object) {
	document := js.Global.Get("document")

	switch element.Type {
	case Normal:
		child := document.Call("createElement", element.Name)

		for _, attr := range element.Attrs {
			child.Call("setAttribute", attr.Name, attr.Value)
		}

		for _, handler := range element.EventHandlers {
			addEventHandler(element, child, &handler)
		}

		for _, elem := range element.Children {
			createElementRecursive(&elem, child)
		}

		domNode.Call("appendChild", child)
	case Text:
		domNode.Set("innerText", element.Attrs[0].Value)
	}
}

func applyPatchToDom(patch *Patch) {
	switch patch.Type {
	case Replace:
		root := js.Global.Get("document").Get("body")

		child := root.Get("lastElementChild")
		if child != nil {
			root.Call("removeChild", child)
		}

		createElementRecursive(&patch.Element, root)
	}
}

//ListenAndServe is a no-op for the javascript target
func ListenAndServe() {}
