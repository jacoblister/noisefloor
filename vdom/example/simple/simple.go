package main

import (
	"github.com/jacoblister/noisefloor/vdom"
)

func onClick(element *vdom.Element, event *vdom.Event) {
	println("click was called from golang callback")
}

func main() {
	e := vdom.NewElement("div",
		"style", "background-color: orange; width: 100; line-height: 6; text-align: center; vertical-align: middle;",
		vdom.NewTextElement("Clicks: 0"),
		vdom.NewEventHandler(vdom.Click, onClick),
	)

	vdom.RenderToDom(e)
}
