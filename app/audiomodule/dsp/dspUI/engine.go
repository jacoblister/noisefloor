package dspUI

import (
	"strconv"
	"strings"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp"
	"github.com/jacoblister/noisefloor/app/vdomcomp"
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

	contextMenu vdomcomp.ContextMenu
}

//MakeEngine create an new Engine Edit UI componenet
func MakeEngine(engine *dsp.Engine, engineState *EngineState) *Engine {
	engineUI := Engine{Engine: engine, state: engineState}
	engine.SetProcessEventFunc(engineUI.processEvent)

	return &engineUI
}

func (e *Engine) processEvent() {
	vdom.UpdateComponentBackground(e)
}

//connectorForProcessor finds the connector give a target
func (e *Engine) connectorForProcessor(processor dsp.Processor, isInput bool, port int) *dsp.Connector {
	for i := 0; i < len(e.Engine.Graph.Connectors); i++ {
		connector := &e.Engine.Graph.Connectors[i]
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
	list := e.Engine.Graph.Connectors
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
			list := e.Engine.Graph.Connectors
			e.Engine.Graph.Connectors = append(list[:index], list[index+1:]...)
		}
		return
	}

	connector.SetProcessor(targetIsInput, targetProcessor)
	connector.SetPort(targetIsInput, targetPort)
	if e.state.editState == connectionAdd {
		e.Engine.Graph.Connectors = append(e.Engine.Graph.Connectors, *connector)
	}
}

// deleteProcessor removes the processor and all its connectors
func (e *Engine) deleteProcessor(processor dsp.Processor) {
	for i := len(e.Engine.Graph.Connectors) - 1; i >= 0; i-- {
		if e.Engine.Graph.Connectors[i].FromProcessor == processor ||
			e.Engine.Graph.Connectors[i].ToProcessor == processor {
			e.Engine.Graph.Connectors = append(e.Engine.Graph.Connectors[:i], e.Engine.Graph.Connectors[i+1:]...)
		}
	}

	for i := len(e.Engine.Graph.Processors) - 1; i >= 0; i-- {
		if e.Engine.Graph.Processors[i].Processor == processor {
			e.Engine.Graph.Processors = append(e.Engine.Graph.Processors[:i], e.Engine.Graph.Processors[i+1:]...)
		}
	}
}

func getUniqueProcessorName(processorName string, existingNames []string) string {
	nameUsed := false
	topIndex := 0
	for i := 0; i < len(existingNames); i++ {
		existingName := existingNames[i]
		if existingName == processorName {
			nameUsed = true
		}
		if strings.HasPrefix(existingName, processorName) && len(existingName) > len(processorName) {
			topIndex = int(existingName[len(processorName)]) - '0'
		}
	}

	if nameUsed {
		processorName += strconv.Itoa(topIndex + 1)
	}
	return processorName
}

// createProcessor adds a new processor at given screen coordinates
func (e *Engine) createProcessor(processorName string, x int, y int) {
	existingNames := []string{}
	for i := 0; i < len(e.Engine.Graph.Processors); i++ {
		existingNames = append(existingNames, e.Engine.Graph.Processors[i].GetName())
	}

	processor := dsp.MakeProcessor(processorName)
	processorDefiniton := dsp.ProcessorDefinition{X: x, Y: y, Processor: processor, Name: getUniqueProcessorName(processorName, existingNames)}
	e.Engine.Graph.Processors = append(e.Engine.Graph.Processors, processorDefiniton)
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
			if event.Type == vdom.ContextMenu {
				// TODO - should not need this, prevent double trigger
				if e.state.contextMenu.Active() {
					return
				}

				x := event.Data["ClientX"].(int)
				y := event.Data["ClientY"].(int)
				e.state.contextMenu = vdomcomp.MakeContextMenu(
					x, y,
					dsp.ListProcessors(), true, func(processorName string) {
						e.createProcessor(processorName, x, y)
						e.Engine.GraphChange(true)
					})
			}
		case ESProcessor:
			processor := event.Data["Processor"].(*dsp.ProcessorDefinition)
			switch event.Type {
			case vdom.MouseDown:
				e.state.selectedProcessor = processor
				e.state.mouseOffsetX = event.Data["OffsetX"].(int) - processor.X
				e.state.mouseOffsetY = event.Data["OffsetY"].(int) - processor.Y
				e.state.editState = moveProcessor
			case vdom.ContextMenu:
				e.state.contextMenu = vdomcomp.MakeContextMenu(
					event.Data["ClientX"].(int),
					event.Data["ClientY"].(int),
					[]string{"Delete"}, true, func(item string) {
						e.deleteProcessor(processor.Processor)
						e.Engine.GraphChange(true)
					})
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
			// snap to grid
			x := event.Data["OffsetX"].(int) - e.state.mouseOffsetX
			y := event.Data["OffsetY"].(int) - e.state.mouseOffsetY
			x = x - (x % procConnWidth)
			y = y - (y % procConnWidth)
			e.state.selectedProcessor.X = x
			e.state.selectedProcessor.Y = y
		case vdom.MouseUp:
			e.Engine.GraphChange(false)
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
				e.Engine.GraphChange(true)
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
				e.Engine.GraphChange(true)
			}
		}
	}
}

func (e *Engine) mainUIEventHandler(element *vdom.Element, event *vdom.Event) {
	event.Data["Source"] = ESMain
	e.handleUIEvent(element, event)
}

// connectorCoordinates returns the coordinates for the connector, which may be being edited
func (e *Engine) connectorCoordinates(
	connector *dsp.Connector,
	fromProcessor *Processor,
	toProcessor *Processor) (x1 int, y1 int, x2 int, y2 int, isConnected bool) {
	procWidth := procDefaultWidth

	if fromProcessor != nil {
		x1, y1 = fromProcessor.GetConnectorPoint(procWidth, false, connector.FromPort)
	}
	if toProcessor != nil {
		x2, y2 = toProcessor.GetConnectorPoint(procWidth, true, connector.ToPort)
	}
	isConnected = true

	if e.state.editState != connectionEdit && e.state.editState != connectionAdd {
		return
	}

	_, targetCount := e.connectorTargetIndex(connector,
		e.state.targetPortIsInput, e.state.targetProcessor, e.state.targetPort)

	// if *connector == *e.state.selectedConnector {	 - TODO simplify once values are out of connector
	if connector.FromProcessor == e.state.selectedConnector.FromProcessor &&
		connector.FromPort == e.state.selectedConnector.FromPort &&
		connector.ToProcessor == e.state.selectedConnector.ToProcessor &&
		connector.ToPort == e.state.selectedConnector.ToPort {
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
			isConnected = false
		}
	}
	return
}

// Render displays the dsp engine frontend.
func (e *Engine) Render() vdom.Element {
	// processors
	processors := []vdom.Component{}
	processorMap := map[dsp.Processor]*Processor{}
	for i := 0; i < len(e.Engine.Graph.Processors); i++ {
		processorDef := &e.Engine.Graph.Processors[i]
		processor := MakeProcessor(processorDef, e.handleUIEvent)
		processors = append(processors, processor)
		processorMap[processorDef.Processor] = processor
	}

	// connectors
	connectionList := e.Engine.Graph.Connectors
	if e.state.editState == connectionAdd {
		connectionList = append(connectionList, *e.state.selectedConnector)
	}
	connectors := []vdom.Component{}
	for i := 0; i < len(connectionList); i++ {
		connector := &connectionList[i]
		x1, y1, x2, y2, isConnected := e.connectorCoordinates(connector, processorMap[connector.FromProcessor], processorMap[connector.ToProcessor])
		connectorLine := &Connector{x1: x1, y1: y1, x2: x2, y2: y2, isConnected: isConnected, value: connector.Value}
		connectors = append(connectors, connectorLine)
	}

	// main view
	elem := vdom.MakeElement("g",
		"id", "dspengineedit",
		vdom.MakeEventHandler(vdom.MouseUp, e.mainUIEventHandler),
		vdom.MakeEventHandler(vdom.MouseDown, e.mainUIEventHandler),
		vdom.MakeEventHandler(vdom.MouseMove, e.mainUIEventHandler),
		vdom.MakeEventHandler(vdom.ContextMenu, e.mainUIEventHandler),
		vdom.MakeElement("rect",
			"x", 0,
			"y", 0,
			"width", 800,
			"height", 600,
			"stroke", "black",
			"fill", "white",
		),
		processors,
		connectors,
		&e.state.contextMenu,
	)

	return elem
}
