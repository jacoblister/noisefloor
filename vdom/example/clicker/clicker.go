package main

import (
	"strconv"

	"github.com/jacoblister/noisefloor/vdom"
)

//Clicker is a click box with a click count
type Clicker struct {
	clicks int
}

func (c *Clicker) onClick(element *vdom.Element, event *vdom.Event) {
	println("click was called from golang callback")

	c.clicks = c.clicks + 1
	vdom.UpdateComponent(c)
}

//Render renders the Clicker component
func (c *Clicker) Render() vdom.Element {
	onClick := func(element *vdom.Element, event *vdom.Event) {
		c.onClick(element, event)
	}

	e := vdom.MakeElement("div",
		"style", "background-color: orange; width: 100; line-height: 6; text-align: center; vertical-align: middle;",
		vdom.MakeElement("span"),
		vdom.MakeTextElement("Clicks: "+strconv.Itoa(c.clicks)),
		vdom.MakeEventHandler(vdom.Click, onClick),
	)
	return e
}

func main() {
	var clicker Clicker

	vdom.RenderComponentToDom(&clicker)
	vdom.ListenAndServe()
}
