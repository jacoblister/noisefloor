package main

import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
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
		"class", "row",
		vdom.MakeElement("div", "class", "col-md-6",
			vdom.MakeElement("div",
				"class", "custom-control custom-checkbox",
				vdom.MakeElement("input",
					"id", prefix+"check",
					"type", "checkbox",
					"class", "custom-control-input",
					checked,
					vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
						item.toggleItem(item.index)
					},
					),
				),
				vdom.MakeElement("label",
					"class", "custom-control-label",
					"for", prefix+"check",
					vdom.MakeTextElement(item.Name),
				),
			),
		),
		vdom.MakeElement("div", "class", "col-md-6",
			vdom.MakeElement("button",
				"id", prefix+"remove",
				"class", "btn btn-outline-primary float-right",
				vdom.MakeTextElement("remove"),
				vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
					item.removeItem(item.index)
				},
				),
			),
		),
	)
	return element
}
