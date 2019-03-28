package main

// Obligatory Todo application

import (
	"github.com/jacoblister/noisefloor/vdom"
)

//Todo is a Todo application with an item list
type Todo struct {
	items []TodoItem
}

//TodoItem is a item with completed status
type TodoItem struct {
	Name      string
	Completed bool
}

func (t *Todo) onChange(element *vdom.Element, event *vdom.Event) {
	t.items = append(t.items, TodoItem{Name: event.Data, Completed: false})
	vdom.UpdateComponent(t)
}

//Render renders the Clicker component
func (t *Todo) Render() vdom.Element {
	onChange := func(element *vdom.Element, event *vdom.Event) {
		t.onChange(element, event)
	}

	items := vdom.MakeElement("div")

	for _, item := range t.items {
		var checked vdom.Attr
		if item.Completed {
			checked = vdom.Attr{Name: "checked", Value: "checked"}
		}

		i := vdom.MakeElement("div",
			vdom.MakeElement("input",
				"type", "checkbox",
				checked,
			),
			vdom.MakeElement("span",
				vdom.MakeTextElement(item.Name),
			),
		)

		items.AppendChild(i)
	}

	result := vdom.MakeElement("div",
		vdom.MakeElement("input",
			vdom.MakeEventHandler(vdom.Change, onChange),
		),
		items,
	)
	return result
}

func main() {
	var todo Todo

	todo.items = append(todo.items, TodoItem{Name: "Implement VDOM", Completed: true})
	todo.items = append(todo.items, TodoItem{Name: "Implement Components", Completed: false})

	vdom.RenderComponentToDom(&todo)
}
