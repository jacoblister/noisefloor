package main

// Simple VDOM iteraction without componenets

import "github.com/jacoblister/noisefloor/vdom"

func onClick(element *vdom.Element, event *vdom.Event) {
	println("click was called from golang callback")
}

func main() {
	e := vdom.MakeElement("div",
		"style", "background-color: orange; width: 100; line-height: 6; text-align: center; vertical-align: middle;",
		vdom.MakeTextElement("Clicks: 0"),
		vdom.MakeEventHandler(vdom.Click, onClick),
	)

	vdom.RenderToDom(e)
}
