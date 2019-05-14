package main

import (
	"github.com/jacoblister/noisefloor/app/vdomcomp"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

//Text is a simple Text component for testing
type Text struct {
	text string
}

//Render renders the Text component
func (t *Text) Render() vdom.Element {
	elem := vdom.MakeElement("text",
		"alignment-baseline", "text-before-edge",
		"x", 0,
		"y", 0,
		vdom.MakeTextElement(t.text),
	)

	return elem
}

//App is a simple SVG example
type App struct {
	hDividerPos    int
	hDividerMoving bool
}

//Render renders the App component
func (a *App) Render() vdom.Element {
	elem := vdom.MakeElement("svg",
		"id", "root",
		"xmlns", "http://www.w3.org/2000/svg",
		"style", "width:100%;height:100%;position:fixed;top:0;left:0;bottom:0;right:0;",
		vdomcomp.MakeLayoutVSplit(640, 480, a.hDividerPos, a.hDividerMoving, &Text{"left"}, &Text{"right"},
			func(pos int) {
				if pos > 100 {
					a.hDividerPos = pos
				}
			},
			func(moving bool) {
				a.hDividerMoving = moving
			},
		),
	)
	return elem
}

func main() {
	app := App{hDividerPos: 320}

	vdom.SetSVGNamespace()
	vdom.RenderComponentToDom(&app)
	vdom.ListenAndServe()
}
