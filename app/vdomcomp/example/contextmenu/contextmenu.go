package main

import (
	"github.com/jacoblister/noisefloor/app/vdomcomp"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

var contextMenu vdomcomp.ContextMenu

//App is a simple SVG example
type App struct{}

//Render renders the App component
func (a *App) Render() vdom.Element {
	text := vdom.MakeElement("text",
		"id", "menu",
		"alignment-baseline", "text-before-edge",
		"x", 200,
		"y", 100,
		vdom.MakeTextElement("menu"),
		vdom.MakeEventHandler(vdom.ContentMenu, func(element *vdom.Element, event *vdom.Event) {
			contextMenu = vdomcomp.MakeContextMenu(
				event.Data["ClientX"].(int),
				event.Data["ClientY"].(int),
				[]string{"First item", "Second item"}, true, func(item string) {
					println("selected:", item)
				})
		}),
	)

	elem := vdom.MakeElement("svg",
		"id", "root",
		"xmlns", "http://www.w3.org/2000/svg",
		"style", "width:100%;height:100%;position:fixed;top:0;left:0;bottom:0;right:0;",
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
