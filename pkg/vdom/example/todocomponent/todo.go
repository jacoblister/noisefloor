package main

// Obligatory Todo application

import (
	"net/http"
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

//Todo is the Todo component
type Todo struct {
	items       []Item
	newItemText string
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
	vdom.UpdateComponent(t)
}

func (t *Todo) removeItem(index int) {
	t.items = append(t.items[:index], t.items[index+1:]...)
	vdom.UpdateComponent(t)
}

func (t *Todo) checkedItemCount() int {
	count := 0
	for i := 0; i < len(t.items); i++ {
		if t.items[i].Completed {
			count++
		}
	}
	return count
}

//Render renders the Clicker component
func (t *Todo) Render() vdom.Element {
	items := vdom.MakeElement("ul", "class", "list-group mb-3")
	for i := 0; i < len(t.items); i++ {
		item := TodoItem{i, t.items[i].Name, t.items[i].Completed,
			func(index int) {
				t.toggleItem(index)
			},
			func(index int) {
				t.removeItem(index)
			},
		}

		li := vdom.MakeElement("li",
			"class", "list-group-item",
			&item)
		items.AppendChild(li)
	}

	result :=
		vdom.MakeElement("div",
			"style", "position: absolute; top: 0; left: 0; height: 100%; width: 100%;",
			"class", "bg-light",
			vdom.MakeElement("div",
				"class", "container",
				vdom.MakeElement("div",
					"class", "py-5 text-center",
					vdom.MakeElement("h2", vdom.MakeTextElement("Todo Check List")),
					vdom.MakeElement("p",
						"class", "lead",
						vdom.MakeTextElement(
							"Add items to the list below, then, the check them off, or remove them ",
						)),
				),
				vdom.MakeElement("div",
					vdom.MakeElement("div",
						"class", "container",
						vdom.MakeElement("h4",
							"class", "row mb-3",
							vdom.MakeTextElement("Todo Items")),
						vdom.MakeElement("div",
							"class", "row",
							vdom.MakeElement("input",
								"id", "addItem",
								"class", "col-md-10 form-control",
								"placeholder", "Add Todo item",
								"value", t.newItemText,
								vdom.MakeEventHandler(vdom.Change, func(element *vdom.Element, event *vdom.Event) {
									t.newItemText = event.Data["Value"].(string)
								}),
							),
							vdom.MakeElement("button",
								"id", "add",
								"class", "col-md-2 btn btn-outline-primary float-right",
								vdom.MakeTextElement("Add Item"),
								vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
									t.addItem(t.newItemText)
									t.newItemText = ""
								}),
							),
						),
					),
					vdom.MakeElement("br"),
					items,
					vdom.MakeElement("br"),
					vdom.MakeElement("div",
						vdom.MakeTextElement("Total items: "+strconv.Itoa(len(t.items))+
							", Checked items: "+strconv.Itoa(t.checkedItemCount())),
					),
				),
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
		vdom.MakeElement("link",
			"rel", "stylesheet",
			"type", "text/css",
			"href", "assets/files/bootstrap.min.css"),
	})

	todo.items = append(todo.items, Item{Name: "Implement VDOM", Completed: true})
	todo.items = append(todo.items, Item{Name: "Implement Components", Completed: false})

	vdom.RenderComponentToDom(&todo)
	// vdom.ListenAndServe("/assets/files/", assets.Assets)
	vdom.ListenAndServe("/assets/files/", http.Dir("assets/files"))
}
