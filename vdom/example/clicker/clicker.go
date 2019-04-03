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
	c.clicks = c.clicks + 1
	vdom.UpdateComponent(c)
}

//Render renders the Clicker component
func (c *Clicker) Render() vdom.Element {
	e := vdom.MakeElement("div",
		"id", "clicker",
		"style", "background-color: orange; width: 100; line-height: 6; text-align: center; vertical-align: middle;",
		vdom.MakeTextElement("Clicks: "+strconv.Itoa(c.clicks)),
		vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
			c.onClick(element, event)
		}),
	)
	return e
}

func main() {
	var clicker Clicker

	vdom.RenderComponentToDom(&clicker)
	vdom.ListenAndServe()
}
