package main

import (
	"strconv"

	"github.com/jacoblister/noisefloor/vdom"
)

// TodoItemToggleFunc is the toggle item callback
type TodoItemToggleFunc func(index int)

// TodoItemRemoveFunc is the remove item callback
type TodoItemRemoveFunc func(index int)

//TodoItem is a item with completed status
type TodoItem struct {
	index     int
	Name      string
	Completed bool

	toggleItem TodoItemToggleFunc
	removeItem TodoItemRemoveFunc
}

//Render renders the TodoItem
func (item *TodoItem) Render() vdom.Element {
	var checked vdom.Attr
	if item.Completed {
		checked = vdom.Attr{Name: "checked", Value: "checked"}
	}

	prefix := strconv.Itoa(item.index) + ":"

	element := vdom.MakeElement("div",
		vdom.MakeElement("input",
			"id", prefix+"check",
			"type", "checkbox",
			checked,
			vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
				item.toggleItem(item.index)
				vdom.UpdateComponent(item)
			},
			),
		),
		vdom.MakeElement("span",
			"id", prefix+"name",
			"style", "display: inline-block; width: 200",
			vdom.MakeTextElement(item.Name),
			vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
				item.toggleItem(item.index)
				vdom.UpdateComponent(item)
			},
			),
		),
		vdom.MakeElement("button",
			"id", prefix+"remove",
			vdom.MakeTextElement("remove"),
			vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
				item.removeItem(item.index)
			},
			),
		),
	)
	return element
}
