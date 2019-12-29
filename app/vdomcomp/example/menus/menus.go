package main

import (
	"github.com/jacoblister/noisefloor/app/vdomcomp"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

var contextMenu vdomcomp.ContextMenu

//App is a simple SVG example
type App struct {
	selected string
}

//Render renders the App component
func (a *App) Render() vdom.Element {
	list := vdomcomp.MakePickList(10, 10, 100, 400, []string{"First", "Second", "Third"}, a.selected, func(item string) {
		a.selected = item
		println("selected:", item)
	})

	text := vdom.MakeElement("text",
		"id", "menu",
		"alignment-baseline", "text-before-edge",
		"x", 200,
		"y", 100,
		vdom.MakeTextElement("menu"),
		vdom.MakeEventHandler(vdom.ContextMenu, func(element *vdom.Element, event *vdom.Event) {
			contextMenu = vdomcomp.MakeContextMenu(
				event.Data["ClientX"].(int),
				event.Data["ClientY"].(int),
				[]string{"First item", "Second item"}, true, func(item string) {
					contextMenu.SetActive(false)
					println("selected:", item)
				})
		}),
	)

	elem := vdom.MakeElement("svg",
		"id", "root",
		"xmlns", "http://www.w3.org/2000/svg",
		"style", "width:100%;height:100%;position:fixed;top:0;left:0;bottom:0;right:0;",
		&list,
		text,
		&contextMenu,
	)
	return elem
}

func main() {
	app := App{}

	vdom.SetSVGNamespace()
	vdom.RenderComponentToDom(&app)
	vdom.ListenAndServe()
}
