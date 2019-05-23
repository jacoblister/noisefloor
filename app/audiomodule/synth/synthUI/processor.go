package synthUI

import (
	"strconv"

	"github.com/jacoblister/noisefloor/app/audiomodule/synth"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

const procWidth = 40
const procHeight = 80
const procConnWidth = 6

// Processor is a synth processor block (with inputs and outputs displayed)
type Processor struct {
	ProcessorDefinition *synth.ProcessorDefinition
	handlerFunc         vdom.HandlerFunc
}

//MakeProcessor create an new Processor UI componenet
func MakeProcessor(processorDefinition *synth.ProcessorDefinition, handlerFunc vdom.HandlerFunc) *Processor {
	processor := Processor{ProcessorDefinition: processorDefinition, handlerFunc: handlerFunc}
	return &processor
}

//GetConnectorPoint gets the co-ordinates of an input or output connection point
func (p *Processor) GetConnectorPoint(isInput bool, port int) (x int, y int) {
	if isInput {
		x = p.ProcessorDefinition.X + (procConnWidth / 2)
		y = p.ProcessorDefinition.Y + (port+1)*(procConnWidth*2) + (procConnWidth / 2)
		return x, y
	}

	x = p.ProcessorDefinition.X + procWidth - (procConnWidth / 2)
	y = p.ProcessorDefinition.Y + (port+1)*(procConnWidth*2) + (procConnWidth / 2)
	return x, y
}

func (p *Processor) processorEventHandler(element *vdom.Element, event *vdom.Event) {
	event.Data["Source"] = ESProcessor
	event.Data["Processor"] = p.ProcessorDefinition
	event.Data["OffsetX"] = event.Data["OffsetX"].(int) + p.ProcessorDefinition.X
	event.Data["OffsetY"] = event.Data["OffsetY"].(int) + p.ProcessorDefinition.Y
	p.handlerFunc(element, event)
}

func (p *Processor) makeConnectorEventHandler(isInput bool, port int) vdom.HandlerFunc {
	return func(element *vdom.Element, event *vdom.Event) {
		event.Data["Source"] = ESConnector
		event.Data["Processor"] = p.ProcessorDefinition
		event.Data["IsInput"] = isInput
		event.Data["Port"] = port
		p.handlerFunc(element, event)
	}
}

// Render displays a processor
func (p *Processor) Render() vdom.Element {
	procName, procInputs, procOutputs := p.ProcessorDefinition.Processor.Definition()

	inConnectors := []vdom.Element{}
	for i := 0; i < len(procInputs); i++ {
		connector := vdom.MakeElement("rect",
			"id", procName+":inconn:"+strconv.Itoa(i),
			"x", p.ProcessorDefinition.X,
			"y", p.ProcessorDefinition.Y+(i+1)*(procConnWidth*2),
			"width", procConnWidth,
			"height", procConnWidth,
			"stroke", "black",
			"fill", "white",
			"cursor", "crosshair",
			vdom.MakeEventHandler(vdom.MouseDown, p.makeConnectorEventHandler(true, i)),
			vdom.MakeEventHandler(vdom.MouseMove, p.makeConnectorEventHandler(true, i)),
		)
		inConnectors = append(inConnectors, connector)
	}

	outConnectors := []vdom.Element{}
	for i := 0; i < len(procOutputs); i++ {
		connector := vdom.MakeElement("rect",
			"id", procName+":outconn:"+strconv.Itoa(i),
			"x", p.ProcessorDefinition.X+procWidth-procConnWidth,
			"y", p.ProcessorDefinition.Y+(i+1)*(procConnWidth*2),
			"width", procConnWidth,
			"height", procConnWidth,
			"stroke", "black",
			"fill", "white",
			"cursor", "crosshair",
			vdom.MakeEventHandler(vdom.MouseDown, p.makeConnectorEventHandler(false, i)),
			vdom.MakeEventHandler(vdom.MouseMove, p.makeConnectorEventHandler(false, i)),
		)
		outConnectors = append(outConnectors, connector)
	}

	element := vdom.MakeElement("g",
		vdom.MakeElement("rect",
			"id", procName,
			"x", p.ProcessorDefinition.X,
			"y", p.ProcessorDefinition.Y,
			"width", procWidth,
			"height", procHeight,
			"stroke", "black",
			"fill", "white",
			"cursor", "pointer",
			vdom.MakeEventHandler(vdom.MouseDown, p.processorEventHandler),
			vdom.MakeEventHandler(vdom.MouseUp, p.processorEventHandler),
			vdom.MakeEventHandler(vdom.MouseMove, p.processorEventHandler),
		),
		inConnectors,
		outConnectors,
	)

	return element
}
