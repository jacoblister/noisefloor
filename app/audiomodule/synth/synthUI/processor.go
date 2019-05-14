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
}

//MakeProcessor create an new Processor UI componenet
func MakeProcessor(processorDefinition *synth.ProcessorDefinition) *Processor {
	processor := Processor{ProcessorDefinition: processorDefinition}
	return &processor
}

//GetConnectorPoint gets to co-ordinates of an input or output connection point
func (p *Processor) GetConnectorPoint(isInput bool, index int) (x int, y int) {
	if isInput {
		x = p.ProcessorDefinition.X + (procConnWidth / 2)
		y = p.ProcessorDefinition.Y + (index+1)*(procConnWidth*2) + (procConnWidth / 2)
		return x, y
	}

	x = p.ProcessorDefinition.X + (procConnWidth / 2)
	y = p.ProcessorDefinition.Y + (index+1)*(procConnWidth*2) + (procConnWidth / 2)
	return x, y
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
		)
		inConnectors = append(inConnectors, connector)
	}

	outConnectors := []vdom.Element{}
	for i := 0; i < len(procOutputs); i++ {
		connector := vdom.MakeElement("rect",
			"id", procName+":inconn:"+strconv.Itoa(i),
			"x", p.ProcessorDefinition.X+procWidth-procConnWidth,
			"y", p.ProcessorDefinition.Y+(i+1)*(procConnWidth*2),
			"width", procConnWidth,
			"height", procConnWidth,
			"stroke", "black",
			"fill", "white",
			"cursor", "crosshair",
		)
		outConnectors = append(outConnectors, connector)
	}

	element := vdom.MakeElement("g",
		vdom.MakeElement("rect",
			"id", "makeosc",
			"x", p.ProcessorDefinition.X,
			"y", p.ProcessorDefinition.Y,
			"width", procWidth,
			"height", procHeight,
			"stroke", "black",
			"fill", "white",
			"cursor", "pointer",
		),
		inConnectors,
		outConnectors,
	)

	return element
}
