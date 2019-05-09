package main

import (
	"strconv"
	"time"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

//Clicker is a click box with a background seconds count
type Clicker struct {
	clicks  int
	seconds int
}

func (c *Clicker) onClick(element *vdom.Element, event *vdom.Event) {
	c.clicks = c.clicks + 1
	vdom.UpdateComponent(c)
}

//Render renders the Clicker component
func (c *Clicker) Render() vdom.Element {
	e := vdom.MakeElement("div",
		"id", "clicker",
		"style", "background-color: orange; width: 100; line-height: 3; text-align: center; vertical-align: middle;",
		vdom.MakeTextElement("Clicks: "+strconv.Itoa(c.clicks)+"\nSeconds: "+strconv.Itoa(c.seconds)),
		vdom.MakeEventHandler(vdom.Click, func(element *vdom.Element, event *vdom.Event) {
			c.onClick(element, event)
		}),
	)
	return e
}

func main() {
	var clicker Clicker

	go func() {
		for {
			time.Sleep(1 * time.Second)
			clicker.seconds++
			vdom.UpdateComponentBackground(&clicker)
		}
	}()

	vdom.RenderComponentToDom(&clicker)
	println("start")
	vdom.ListenAndServe()
	println("served")
}
