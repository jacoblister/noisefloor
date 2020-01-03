package dspUI

import (
	"strconv"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

const procDefaultWidth = 80
const procDefaultHeight = 80
const procConnWidth = 8

var level int

// Processor is a dsp processor block (with inputs and outputs displayed)
type Processor struct {
	ProcessorDefinition *processor.Definition
	handlerFunc         vdom.HandlerFunc
}

//MakeProcessor create an new Processor UI componenet
func MakeProcessor(processorDefinition *processor.Definition, handlerFunc vdom.HandlerFunc) *Processor {
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

func (p *Processor) makeParameterEventHandler(index int, factor float32) vdom.HandlerFunc {
	return func(element *vdom.Element, event *vdom.Event) {
		value := float32(event.Data["OffsetX"].(int)) / factor
		p.ProcessorDefinition.Processor.SetParameter(index, value)
	}
}

// Render displays a processor
func (p *Processor) Render() vdom.Element {
	maxConnectors := p.ProcessorDefinition.MaxConnectors()
	procName := p.ProcessorDefinition.Name
	procDefName, procInputs, procOutputs, procParameters := p.ProcessorDefinition.Processor.Definition()
	if len(procName) == 0 {
		procName = procDefName
	}

	procWidth := procDefaultWidth
	procHeight := (maxConnectors + 2 + len(procParameters)) * procConnWidth * 2

	customRenderDimentions, ok := p.ProcessorDefinition.Processor.(customRenderDimentions)
	if ok {
		procWidth, procHeight = customRenderDimentions.CustomRenderDimentions()
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

	parameters := []vdom.Element{}
	for i := 0; i < len(procParameters); i++ {
		_, y := p.GetConnectorPoint(procWidth, true, i+maxConnectors)

		width := procWidth - ((procConnWidth + 2) * 2)
		levelValue := strconv.FormatFloat(float64(procParameters[i].Value), 'f', -1, 32)
		levelWidth := int(float32(width) * (procParameters[i].Value / procParameters[i].Max))
		levelFactor := float32(width) / procParameters[i].Max

		level := vdom.MakeElement("rect",
			"id", procName+":parameter:"+strconv.Itoa(i),
			"x", p.ProcessorDefinition.X+procConnWidth+2,
			"y", y-2,
			"width", levelWidth,
			"height", procConnWidth+4,
			"stroke", "none",
			"fill", "cyan",
			"pointer-events", "none",
		)
		parameters = append(parameters, level)

		bound := vdom.MakeElement("rect",
			"id", procName+":parameter:"+strconv.Itoa(i),
			"x", p.ProcessorDefinition.X+procConnWidth+2,
			"y", y-2,
			"width", width,
			"height", procConnWidth+4,
			"stroke", "black",
			"fill", "none",
			"pointer-events", "all",
			"cursor", "crosshair",
			vdom.MakeEventHandler(vdom.MouseDown, p.makeParameterEventHandler(i, levelFactor)),
			vdom.MakeElement("title", vdom.MakeTextElement(levelValue)),
		)
		parameters = append(parameters, bound)

		name := vdom.MakeElement("text",
			"font-family", "sans-serif",
			"text-anchor", "middle",
			"alignment-baseline", "hanging",
			"font-size", 10,
			"x", p.ProcessorDefinition.X+(procWidth/2),
			"y", y,
			vdom.MakeTextElement(procParameters[i].Name))
		parameters = append(parameters, name)
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
			vdom.MakeEventHandler(vdom.ContextMenu, p.processorEventHandler),
		),
		procNameLabel,
		procLine,
		inConnectors,
		outConnectors,
		parameters,
	)

	custom, ok := p.ProcessorDefinition.Processor.(vdom.Component)
	if ok {
		element.AppendChild(vdom.MakeElement("g",
			"transform", "translate("+strconv.Itoa(p.ProcessorDefinition.X)+","+strconv.Itoa(p.ProcessorDefinition.Y)+")",
			custom,
		))
	}

	return element
}
