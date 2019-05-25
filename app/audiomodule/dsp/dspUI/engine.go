package dspUI

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

// editState enumerated type
type editState int

// editing state
const (
	idle editState = iota
	moveProcessor
	connectionEdit
	connectionAdd
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

// Engine is the dsp engine UI
type Engine struct {
	Engine *dsp.Engine
	state  *EngineState
}

// EngineState is the dsp engine UI stateful store
type EngineState struct {
	editState                editState
	selectedProcessor        *dsp.ProcessorDefinition
	selectedConnector        *dsp.Connector
	selectedConnectorIsInput bool
	targetProcessor          dsp.Processor
	targetPort               int
	targetPortIsInput        bool
	mouseOffsetX             int
	mouseOffsetY             int
	mouseIgnore              bool
}

//MakeEngine create an new Engine Edit UI componenet
func MakeEngine(engine *dsp.Engine, engineState *EngineState) *Engine {
	engineUI := Engine{Engine: engine, state: engineState}
	return &engineUI
}

//connectorForProcessor finds the connector give a target
func (e *Engine) connectorForProcessor(processor dsp.Processor, isInput bool, port int) *dsp.Connector {
	for i := 0; i < len(e.Engine.Graph.ConnectorList); i++ {
		connector := &e.Engine.Graph.ConnectorList[i]
		if connector.Processor(isInput) == processor && connector.Port(isInput) == port {
			return connector
		}
	}
	return nil
}

// connectorTargetIndex iterates the connector list,
// and gets the index and count of current connections at target
func (e *Engine) connectorTargetIndex(connector *dsp.Connector,
	targetIsInput bool, targetProcessor dsp.Processor, targetPort int) (index int, count int) {
	count = 0
	index = -1
	list := e.Engine.Graph.ConnectorList
	for i := 0; i < len(list); i++ {
		if &list[i] == connector {
			index = i
			continue
		}
		if list[i].Processor(targetIsInput) == targetProcessor &&
			list[i].Port(targetIsInput) == targetPort {
			count++
		}
	}

	return index, count
}

// updateConnector updates the connector list after change (create, modify, delete)
func (e *Engine) updateConnector(connector *dsp.Connector,
	targetIsInput bool, targetProcessor dsp.Processor, targetPort int) {

	index, targetCount := e.connectorTargetIndex(connector, targetIsInput, targetProcessor, targetPort)

	// allow input to output connections only
	// do not allow connection if connector already exists on input connector
	if e.state.selectedConnectorIsInput != targetIsInput ||
		targetIsInput && targetCount > 0 {
		targetProcessor = nil
	}

	// delete connector operation
	if targetProcessor == nil {
		// only update if not in add
		if e.state.editState == connectionEdit {
			list := e.Engine.Graph.ConnectorList
			e.Engine.Graph.ConnectorList = append(list[:index], list[index+1:]...)
		}
		return
	}

	connector.SetProcessor(targetIsInput, targetProcessor)
	connector.SetPort(targetIsInput, targetPort)
	if e.state.editState == connectionAdd {
		e.Engine.Graph.ConnectorList = append(e.Engine.Graph.ConnectorList, *connector)
	}
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
				processor := event.Data["Processor"].(*dsp.ProcessorDefinition)
				e.state.selectedProcessor = processor
				e.state.mouseOffsetX = event.Data["OffsetX"].(int) - processor.X
				e.state.mouseOffsetY = event.Data["OffsetY"].(int) - processor.Y
				e.state.editState = moveProcessor
			}
		case ESConnector:
			switch event.Type {
			case vdom.MouseDown:
				isInput := event.Data["IsInput"].(bool)
				port := event.Data["Port"].(int)
				processor := event.Data["Processor"].(*dsp.ProcessorDefinition)
				connector := e.connectorForProcessor(processor.Processor, isInput, port)

				e.state.selectedConnectorIsInput = isInput
				if connector == nil {
					// new connector
					connector = &dsp.Connector{}
					connector.SetProcessor(isInput, processor.Processor)
					connector.SetPort(isInput, port)
					e.state.selectedConnectorIsInput = !isInput
					e.state.editState = connectionAdd
				} else {
					// existing connector
					e.state.editState = connectionEdit
				}
				e.state.selectedProcessor = processor
				e.state.selectedConnector = connector
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
	case connectionEdit:
		fallthrough
	case connectionAdd:
		switch event.Data["Source"] {
		case ESMain:
			switch event.Type {
			case vdom.MouseMove:
				if !e.state.mouseIgnore {
					// ignore if previous event was mouse move (repeat from connector/processor)
					e.state.mouseOffsetX = event.Data["OffsetX"].(int)
					e.state.mouseOffsetY = event.Data["OffsetY"].(int)
					e.state.targetProcessor = nil
				}
				e.state.mouseIgnore = false
			case vdom.MouseUp:
				e.updateConnector(e.state.selectedConnector, e.state.targetPortIsInput,
					e.state.targetProcessor, e.state.targetPort)
				e.state.editState = idle
			}
		case ESConnector:
			switch event.Type {
			case vdom.MouseMove:
				processor := event.Data["Processor"].(*dsp.ProcessorDefinition)
				e.state.targetPortIsInput = event.Data["IsInput"].(bool)
				e.state.targetProcessor = processor.Processor
				e.state.targetPort = event.Data["Port"].(int)
				e.state.mouseIgnore = true
			case vdom.MouseUp:
				e.updateConnector(e.state.selectedConnector, e.state.targetPortIsInput,
					e.state.targetProcessor, e.state.targetPort)
				e.state.editState = idle
			}
		}
	}
}

func (e *Engine) mainUIEventHandler(element *vdom.Element, event *vdom.Event) {
	event.Data["Source"] = ESMain
	e.handleUIEvent(element, event)
}

// connectorCoordinates returns the coordinates for the connector, which may be being edited
func (e *Engine) connectorCoordinates(connector *dsp.Connector, fromProcessor *Processor, toProcessor *Processor) (x1 int, y1 int, x2 int, y2 int, stroke string) {
	if fromProcessor != nil {
		x1, y1 = fromProcessor.GetConnectorPoint(false, connector.FromPort)
	}
	if toProcessor != nil {
		x2, y2 = toProcessor.GetConnectorPoint(true, connector.ToPort)
	}
	stroke = "darkblue"

	if e.state.editState != connectionEdit && e.state.editState != connectionAdd {
		return
	}

	_, targetCount := e.connectorTargetIndex(connector,
		e.state.targetPortIsInput, e.state.targetProcessor, e.state.targetPort)

	if *connector == *e.state.selectedConnector {
		if e.state.selectedConnectorIsInput {
			x2 = e.state.mouseOffsetX
			y2 = e.state.mouseOffsetY
		} else {
			x1 = e.state.mouseOffsetX
			y1 = e.state.mouseOffsetY
		}
		if e.state.targetProcessor == nil ||
			e.state.targetPortIsInput != e.state.selectedConnectorIsInput ||
			(e.state.targetPortIsInput && targetCount > 0) {
			stroke = "grey"
		}
	}
	return
}

// Render displays the dsp engine frontend.
func (e *Engine) Render() vdom.Element {
	// processors
	processors := []vdom.Component{}
	processorMap := map[dsp.Processor]*Processor{}
	for i := 0; i < len(e.Engine.Graph.ProcessorList); i++ {
		processorDef := &e.Engine.Graph.ProcessorList[i]
		processor := MakeProcessor(processorDef, e.handleUIEvent)
		processors = append(processors, processor)
		processorMap[processorDef.Processor] = processor
	}

	// connectors
	connectionList := e.Engine.Graph.ConnectorList
	if e.state.editState == connectionAdd {
		connectionList = append(connectionList, *e.state.selectedConnector)
	}
	connectors := []vdom.Element{}
	for i := 0; i < len(connectionList); i++ {
		connector := &connectionList[i]
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

	// main view
	elem := vdom.MakeElement("g",
		"id", "dspengineedit",
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
