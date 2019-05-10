package synthUI

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/synth"
	"github.com/jacoblister/noisefloor/pkg/vdom"
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

// Engine is the synth engine UI
type Engine struct {
	Engine *synth.Engine

	editState editState
}

//MakeEngine create an new Engine Edit UI componenet
func MakeEngine(engine *synth.Engine) *Engine {
	engineUI := Engine{Engine: engine}
	return &engineUI
}

// Render displays the synth engine frontend,
// based on the current editing state
func (e *Engine) handleUIEvent() {
}

// Render displays the synth engine frontend.
func (e *Engine) Render() vdom.Element {
	processors := []vdom.Component{}
	for i := 0; i < len(e.Engine.Graph.ProcessorList); i++ {
		processor := e.Engine.Graph.ProcessorList[i]
		processors = append(processors, MakeProcessor(&processor))
	}

	elem := vdom.MakeElement("g",
		"id", "synthengineedit",
		vdom.MakeEventHandler(vdom.MouseMove, func(element *vdom.Element, event *vdom.Event) {
			println("mouse move x=", event.Data["OffsetX"].(int), " y=", event.Data["OffsetY"].(int))
		}),
		vdom.MakeElement("rect",
			"x", "0",
			"y", "0",
			"width", "640",
			"height", "480",
			"stroke", "black",
			"fill", "white",
		),
		processors,
	)

	return elem
}
