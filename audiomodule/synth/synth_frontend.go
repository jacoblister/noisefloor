package synth

import (
	"strconv"

	"github.com/jacoblister/noisefloor/vdom"
)

// editState enumerated type
type editState int

// editing state
const (
	idle editState = iota
	placeProcessor
	connectNode
	selectItems
)

type engineFrontend struct {
	editState editState
}

// Render displays the synth engine frontend,
// based on the current editing state
func (e *Engine) handleUIEvent() {
}

// Render displays the synth engine frontend.
func (e *Engine) Render() vdom.Element {
	processors := []vdom.Element{}
	for i := 0; i < len(e.processorGraph.processorList); i++ {
		processor := e.processorGraph.processorList[i]
		processors = append(processors,
			vdom.MakeElement("rect",
				"id", "makeosc",
				"x", strconv.Itoa(processor.x),
				"y", strconv.Itoa(processor.y),
				"width", "40",
				"height", "20",
				"stroke", "black",
				"fill", "white",
			),
		)
	}

	elem := vdom.MakeElement("g",
		"id", "synthengineedit",
		vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
			println("mouse move x=", event.Data["ClientX"].(int), " y=", event.Data["ClientY"].(int))
		}),
		vdom.MakeElement("rect",
			"x", "100",
			"y", "100",
			"width", "640",
			"height", "480",
			"stroke", "black",
			"fill", "white",
		),
		processors,
		// vdom.MakeElement("rect",
		// 	"id", "makeosc",
		// 	"x", "110",
		// 	"y", "110",
		// 	"width", "40",
		// 	"height", "20",
		// 	"stroke", "black",
		// 	"fill", "white",
		// 	vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
		// 		println("osc move x=", event.Data["ClientX"].(int), " y=", event.Data["ClientY"].(int))
		// 	}),
		//		),
	)

	return elem
}
