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
	moveProcessor
	connectNodes
	selectItems
)

// Engine is the synth engine UI
type Engine struct {
	Engine *synth.Engine
	state  *EngineState
}

// EngineState is the synth engine UI stateful store
type EngineState struct {
	editState         editState
	selectedProcessor *synth.ProcessorDefinition
	mouseOffsetX      int
	mouseOffsetY      int
}

//MakeEngine create an new Engine Edit UI componenet
func MakeEngine(engine *synth.Engine, engineState *EngineState) *Engine {
	engineUI := Engine{Engine: engine, state: engineState}
	return &engineUI
}

// handleUIEvent processes a User Interface event,
// based on the current editing state
func (e *Engine) handleUIEvent(element *vdom.Element, event *vdom.Event) {
	switch e.state.editState {
	case idle:
		switch event.Data["Source"] {
		case "processor":
			switch event.Type {
			case "mousedown":
				processor := event.Data["Processor"].(*synth.ProcessorDefinition)
				e.state.selectedProcessor = processor
				e.state.mouseOffsetX = event.Data["OffsetX"].(int) - processor.X
				e.state.mouseOffsetY = event.Data["OffsetY"].(int) - processor.Y
				e.state.editState = moveProcessor
			}
		}
	case moveProcessor:
		switch event.Type {
		case "mousemove":
			e.state.selectedProcessor.X = event.Data["OffsetX"].(int) - e.state.mouseOffsetX
			e.state.selectedProcessor.Y = event.Data["OffsetY"].(int) - e.state.mouseOffsetY
		case "mouseup":
			e.state.editState = idle
		}
	}
}

func (e *Engine) mainUIEventHandler(element *vdom.Element, event *vdom.Event) {
	event.Data["Source"] = "main"
	e.handleUIEvent(element, event)
}

// Render displays the synth engine frontend.
func (e *Engine) Render() vdom.Element {
	processors := []vdom.Component{}
	processorMap := map[synth.Processor]*Processor{}
	for i := 0; i < len(e.Engine.Graph.ProcessorList); i++ {
		processorDef := &e.Engine.Graph.ProcessorList[i]
		processor := MakeProcessor(processorDef, e.handleUIEvent)
		processors = append(processors, processor)
		processorMap[processorDef.Processor] = processor
	}

	connectors := []vdom.Element{}
	for i := 0; i < len(e.Engine.Graph.ConnectorList); i++ {
		connector := e.Engine.Graph.ConnectorList[i]
		x1, y1 := processorMap[connector.FromProcessor].GetConnectorPoint(false, connector.FromPort)
		x2, y2 := processorMap[connector.ToProcessor].GetConnectorPoint(true, connector.ToPort)
		line := vdom.MakeElement("line",
			"x1", float64(x1)+0.5,
			"y1", float64(y1)+0.5,
			"x2", float64(x2)+0.5,
			"y2", float64(y2)+0.5,
			"stroke", "darkblue",
		)
		connectors = append(connectors, line)
	}

	elem := vdom.MakeElement("g",
		"id", "synthengineedit",
		vdom.MakeEventHandler(vdom.MouseUp, e.mainUIEventHandler),
		vdom.MakeEventHandler(vdom.MouseDown, e.mainUIEventHandler),
		vdom.MakeEventHandler(vdom.MouseMove, e.mainUIEventHandler),
		vdom.MakeElement("rect",
			"x", "0",
			"y", "0",
			"width", "640",
			"height", "480",
			"stroke", "black",
			"fill", "white",
		),
		connectors,
		processors,
	)

	return elem
}
