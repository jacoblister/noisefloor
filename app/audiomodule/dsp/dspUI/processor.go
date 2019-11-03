package dspUI

import (
	"strconv"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

const procDefaultWidth = 80
const procDefaultHeight = 80
const procConnWidth = 8

var level int

// Processor is a dsp processor block (with inputs and outputs displayed)
type Processor struct {
	ProcessorDefinition *dsp.ProcessorDefinition
	handlerFunc         vdom.HandlerFunc
}

//MakeProcessor create an new Processor UI componenet
func MakeProcessor(processorDefinition *dsp.ProcessorDefinition, handlerFunc vdom.HandlerFunc) *Processor {
	processor := Processor{ProcessorDefinition: processorDefinition, handlerFunc: handlerFunc}
	return &processor
}

type customRenderDimentions interface {
	CustomRenderDimentions() (width int, height int)
}

//GetConnectorPoint gets the mid point co-ordinates of an input or output connection point
func (p *Processor) GetConnectorPoint(procWidth int, isInput bool, port int) (x int, y int) {
	y = p.ProcessorDefinition.Y + port*(procConnWidth*2) + (procConnWidth * 4)

	if isInput {
		x = p.ProcessorDefinition.X + (procConnWidth / 2)
		return x, y
	}

	x = p.ProcessorDefinition.X + procWidth - (procConnWidth / 2)
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
	procWidth := procDefaultWidth
	procHeight := (p.ProcessorDefinition.MaxConnectors() + 2) * procConnWidth * 2
	customRenderDimentions, ok := p.ProcessorDefinition.Processor.(customRenderDimentions)
	if ok {
		procWidth, procHeight = customRenderDimentions.CustomRenderDimentions()
	}

	procName, procInputs, procOutputs := p.ProcessorDefinition.Processor.Definition()
	inConnectors := []vdom.Element{}
	for i := 0; i < len(procInputs); i++ {
		x, y := p.GetConnectorPoint(procWidth, true, i)
		connector := vdom.MakeElement("rect",
			"id", procName+":inconn:"+strconv.Itoa(i),
			"x", x-procConnWidth/2,
			"y", y-procConnWidth/2,
			"width", procConnWidth,
			"height", procConnWidth,
			"stroke", "black",
			"fill", "none",
			"pointer-events", "all",
			"cursor", "crosshair",
			vdom.MakeEventHandler(vdom.MouseDown, p.makeConnectorEventHandler(true, i)),
			vdom.MakeEventHandler(vdom.MouseMove, p.makeConnectorEventHandler(true, i)),
		)
		inConnectors = append(inConnectors, connector)

		label := vdom.MakeElement("text",
			"font-family", "sans-serif",
			"text-anchor", "start",
			"alignment-baseline", "middle",
			"font-size", 10,
			"x", x+(procConnWidth/2+2),
			"y", y,
			vdom.MakeTextElement(procInputs[i]),
		)
		inConnectors = append(inConnectors, label)
	}

	outConnectors := []vdom.Element{}
	for i := 0; i < len(procOutputs); i++ {
		x, y := p.GetConnectorPoint(procWidth, false, i)
		connector := vdom.MakeElement("rect",
			"id", procName+":outconn:"+strconv.Itoa(i),
			"x", x-procConnWidth/2,
			"y", y-procConnWidth/2,
			"width", procConnWidth,
			"height", procConnWidth,
			"stroke", "black",
			"fill", "none",
			"pointer-events", "all",
			"cursor", "crosshair",
			vdom.MakeEventHandler(vdom.MouseDown, p.makeConnectorEventHandler(false, i)),
			vdom.MakeEventHandler(vdom.MouseMove, p.makeConnectorEventHandler(false, i)),
		)
		outConnectors = append(outConnectors, connector)

		label := vdom.MakeElement("text",
			"font-family", "sans-serif",
			"text-anchor", "end",
			"alignment-baseline", "middle",
			"font-size", 10,
			"x", x-(procConnWidth/2+4),
			"y", y,
			vdom.MakeTextElement(procOutputs[i]),
		)
		outConnectors = append(outConnectors, label)
	}

	procNameLabel := vdom.MakeElement("text",
		"font-family", "sans-serif",
		"text-anchor", "middle",
		"alignment-baseline", "hanging",
		"font-size", 10,
		"x", p.ProcessorDefinition.X+procWidth/2,
		"y", p.ProcessorDefinition.Y+4,
		vdom.MakeTextElement(procName),
	)
	procLine := vdom.MakeElement("line",
		"stroke", "black",
		"x1", float64(p.ProcessorDefinition.X)+0.5,
		"y1", float64(p.ProcessorDefinition.Y)+16+0.5,
		"x2", float64(p.ProcessorDefinition.X)+float64(procWidth)+0.5,
		"y2", float64(p.ProcessorDefinition.Y)+16+0.5,
	)

	component, _ := p.ProcessorDefinition.Processor.(vdom.Component)

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
		procNameLabel,
		procLine,
		inConnectors,
		outConnectors,
		vdom.MakeElement("g",
			"transform", "translate("+strconv.Itoa(p.ProcessorDefinition.X)+","+strconv.Itoa(p.ProcessorDefinition.Y)+")",
			component,
		),
	)

	return element
}
