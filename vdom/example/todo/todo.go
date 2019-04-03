package main

// Obligatory Todo application

import (
	"strconv"

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

func (t *Todo) addItem(name string) {
	t.items = append(t.items, TodoItem{Name: name, Completed: false})
	vdom.UpdateComponent(t)
}

func (t *Todo) toggleItem(item *TodoItem) {
	item.Completed = !item.Completed
	vdom.UpdateComponent(t)
}

func (t *Todo) removeItem(item *TodoItem) {
	for i := 0; i < len(t.items); i++ {
		if &t.items[i] == item {
			t.items = append(t.items[:i], t.items[i+1:]...)
			vdom.UpdateComponent(t)
			return
		}
	}
}

func (t *Todo) renderItem(item *TodoItem, index int) vdom.Element {
	var checked vdom.Attr
	if item.Completed {
		checked = vdom.Attr{Name: "checked", Value: "checked"}
	}

	prefix := strconv.Itoa(index) + ":"

	element := vdom.MakeElement("div",
		vdom.MakeElement("input",
			"id", prefix+"check",
			"type", "checkbox",
			checked,
			vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
				t.toggleItem(item)
			},
			),
		),
		vdom.MakeElement("span",
			"id", prefix+"name",
			"style", "display: inline-block; width: 200",
			vdom.MakeTextElement(item.Name),
			vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
				t.toggleItem(item)
			},
			),
		),
		vdom.MakeElement("button",
			"id", prefix+"remove",
			vdom.MakeTextElement("remove"),
			vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
				t.removeItem(item)
			},
			),
		),
	)
	return element
}

//Render renders the Clicker component
func (t *Todo) Render() vdom.Element {
	items := vdom.MakeElement("div")

	for i := 0; i < len(t.items); i++ {
		element := t.renderItem(&t.items[i], i)
		items.AppendChild(element)
	}

	result := vdom.MakeElement("div",
		vdom.MakeElement("input",
			"id", "addItem",
			"placeholder", "add TODO item",
			vdom.MakeEventHandler(vdom.Change, func(element *vdom.Element, event *vdom.Event) {
				t.addItem(event.Data)
			},
			),
		),
		items,
		vdom.MakeElement("br"),
		vdom.MakeElement("div",
			vdom.MakeTextElement("Total items: "+strconv.Itoa(len(t.items))),
		),
	)
	return result
}

func main() {
	var todo Todo

	todo.items = append(todo.items, TodoItem{Name: "Implement VDOM", Completed: true})
	todo.items = append(todo.items, TodoItem{Name: "Implement Components", Completed: false})

	vdom.RenderComponentToDom(&todo)
	vdom.ListenAndServe()
}
