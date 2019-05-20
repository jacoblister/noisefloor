package synthUI

import (
	"fmt"

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

// eventSource is the enttity from where to event originated
type eventSource int

// event source
const (
	ESMain eventSource = iota
	ESProcessor
	ESConnector
)

// Engine is the synth engine UI
type Engine struct {
	Engine *synth.Engine
	state  *EngineState
}

// EngineState is the synth engine UI stateful store
type EngineState struct {
	editState                editState
	selectedProcessor        *synth.ProcessorDefinition
	selectedConnector        *synth.Connector
	selectedConnectorIsInput bool
	mouseOffsetX             int
	mouseOffsetY             int
}

//MakeEngine create an new Engine Edit UI componenet
func MakeEngine(engine *synth.Engine, engineState *EngineState) *Engine {
	engineUI := Engine{Engine: engine, state: engineState}
	return &engineUI
}

func (e *Engine) getConnectorForProcessor(processor synth.Processor, isInput bool, index int) *synth.Connector {
	for i := 0; i < len(e.Engine.Graph.ConnectorList); i++ {
		connector := &e.Engine.Graph.ConnectorList[i]
		if isInput && connector.ToProcessor == processor && connector.ToPort == index {
			return connector
		}
		if !isInput && connector.FromProcessor == processor && connector.FromPort == index {
			return connector
		}
	}
	return nil
}

// handleUIEvent processes a User Interface event,
// based on the current editing state
func (e *Engine) handleUIEvent(element *vdom.Element, event *vdom.Event) {
	switch e.state.editState {
	case idle:
		switch event.Data["Source"] {
		case ESMain:
			e.state.mouseOffsetX = event.Data["OffsetX"].(int)
			e.state.mouseOffsetY = event.Data["OffsetY"].(int)
		case ESProcessor:
			switch event.Type {
			case vdom.MouseDown:
				processor := event.Data["Processor"].(*synth.ProcessorDefinition)
				e.state.selectedProcessor = processor
				e.state.mouseOffsetX = event.Data["OffsetX"].(int) - processor.X
				e.state.mouseOffsetY = event.Data["OffsetY"].(int) - processor.Y
				e.state.editState = moveProcessor
			}
		case ESConnector:
			switch event.Type {
			case vdom.MouseDown:
				processor := event.Data["Processor"].(*synth.ProcessorDefinition)
				connector := e.getConnectorForProcessor(
					processor.Processor,
					event.Data["IsInput"].(bool),
					event.Data["Index"].(int))

				e.state.selectedProcessor = processor
				e.state.selectedConnector = connector
				e.state.selectedConnectorIsInput = event.Data["IsInput"].(bool)
				e.state.editState = connectNodes
			}
		}
	case moveProcessor:
		switch event.Type {
		case vdom.MouseMove:
			e.state.selectedProcessor.X = event.Data["OffsetX"].(int) - e.state.mouseOffsetX
			e.state.selectedProcessor.Y = event.Data["OffsetY"].(int) - e.state.mouseOffsetY
		case vdom.MouseUp:
			e.state.editState = idle
		}
	case connectNodes:
		switch event.Data["Source"] {
		case ESMain:
			switch event.Type {
			case vdom.MouseMove:
				e.state.mouseOffsetX = event.Data["OffsetX"].(int)
				e.state.mouseOffsetY = event.Data["OffsetY"].(int)
			case vdom.MouseUp:
				println("connect exit")
				e.state.editState = idle
			}
		case ESConnector:
			switch event.Type {
			case vdom.MouseMove:
				processor := event.Data["Processor"].(*synth.ProcessorDefinition)
				fmt.Println(e.state.selectedConnectorIsInput)
				if e.state.selectedConnectorIsInput {
					e.state.selectedConnector.ToProcessor = processor.Processor
					e.state.selectedConnector.ToPort = event.Data["Index"].(int)
				} else {
					e.state.selectedConnector.FromProcessor = processor.Processor
					e.state.selectedConnector.FromPort = event.Data["Index"].(int)
				}
				println("connector plugged")
			case vdom.MouseUp:
				println("connector plugged")
			}
		}
	}
}

func (e *Engine) mainUIEventHandler(element *vdom.Element, event *vdom.Event) {
	event.Data["Source"] = ESMain
	e.handleUIEvent(element, event)
}

// connectorCoordinates returns the coordinates for the connector, which may be being edited
func (e *Engine) connectorCoordinates(connector *synth.Connector, fromProcessor *Processor, toProcessor *Processor) (x1 int, y1 int, x2 int, y2 int, stroke string) {
	x1, y1 = fromProcessor.GetConnectorPoint(false, connector.FromPort)
	x2, y2 = toProcessor.GetConnectorPoint(true, connector.ToPort)
	stroke = "darkblue"

	if e.state.editState == connectNodes {
		if connector == e.state.selectedConnector {
			if e.state.selectedConnectorIsInput {
				x2 = e.state.mouseOffsetX
				y2 = e.state.mouseOffsetY
			} else {
				x1 = e.state.mouseOffsetX
				y1 = e.state.mouseOffsetY
			}
			stroke = "grey"
		}
	}

	return
}

// Render displays the synth engine frontend.
func (e *Engine) Render() vdom.Element {
	// processors
	processors := []vdom.Component{}
	processorMap := map[synth.Processor]*Processor{}
	for i := 0; i < len(e.Engine.Graph.ProcessorList); i++ {
		processorDef := &e.Engine.Graph.ProcessorList[i]
		processor := MakeProcessor(processorDef, e.handleUIEvent)
		processors = append(processors, processor)
		processorMap[processorDef.Processor] = processor
	}

	// connectors
	connectors := []vdom.Element{}
	for i := 0; i < len(e.Engine.Graph.ConnectorList); i++ {
		connector := &e.Engine.Graph.ConnectorList[i]
		x1, y1, x2, y2, stroke := e.connectorCoordinates(connector, processorMap[connector.FromProcessor], processorMap[connector.ToProcessor])
		line := vdom.MakeElement("line",
			"x1", float64(x1)+0.5,
			"y1", float64(y1)+0.5,
			"x2", float64(x2)+0.5,
			"y2", float64(y2)+0.5,
			"stroke", stroke,
		)
		connectors = append(connectors, line)
	}

	// maim view
	elem := vdom.MakeElement("g",
		"id", "synthengineedit",
		vdom.MakeEventHandler(vdom.MouseUp, e.mainUIEventHandler),
		vdom.MakeEventHandler(vdom.MouseDown, e.mainUIEventHandler),
		vdom.MakeEventHandler(vdom.MouseMove, e.mainUIEventHandler),
		vdom.MakeElement("rect",
			"x", 0,
			"y", 0,
			"width", 640,
			"height", 400,
			"stroke", "black",
			"fill", "white",
		),
		connectors,
		processors,
	)

	return elem
}
