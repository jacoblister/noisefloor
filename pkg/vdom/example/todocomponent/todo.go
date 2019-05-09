package main

// Obligatory Todo application

import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
	"github.com/jacoblister/noisefloor/pkg/vdom/example/todocomponent/assets"
)

//Todo is the Todo component
type Todo struct {
	items []Item
}

//Item is the todo item state
type Item struct {
	Name      string
	Completed bool
}

func (t *Todo) addItem(name string) {
	t.items = append(t.items, Item{Name: name, Completed: false})
	vdom.UpdateComponent(t)
}

func (t *Todo) toggleItem(index int) {
	t.items[index].Completed = !t.items[index].Completed
	// vdom.UpdateComponent(t)
}

func (t *Todo) removeItem(index int) {
	t.items = append(t.items[:index], t.items[index+1:]...)
	vdom.UpdateComponent(t)
}

//Render renders the Clicker component
func (t *Todo) Render() vdom.Element {
	items := []vdom.Component{}

	for i := 0; i < len(t.items); i++ {
		item := TodoItem{i, t.items[i].Name, t.items[i].Completed,
			func(index int) {
				t.toggleItem(index)
			},
			func(index int) {
				t.removeItem(index)
			},
		}
		items = append(items, &item)
	}

	result := vdom.MakeElement("div",
		vdom.MakeElement("input",
			"id", "addItem",
			"placeholder", "add TODO item",
			vdom.MakeEventHandler(vdom.Change, func(element *vdom.Element, event *vdom.Event) {
				t.addItem(event.Data["Value"].(string))
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

	vdom.SetHeaderElements([]vdom.Element{
		vdom.MakeElement("link",
			"rel", "stylesheet",
			"type", "text/css",
			"href", "assets/files/style.css"),
	})

	todo.items = append(todo.items, Item{Name: "Implement VDOM", Completed: true})
	todo.items = append(todo.items, Item{Name: "Implement Components", Completed: false})

	vdom.RenderComponentToDom(&todo)
	vdom.ListenAndServe("/assets/files/", assets.Assets)
}
