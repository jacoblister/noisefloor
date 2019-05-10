package synthUI

import (
	"strconv"

	"github.com/jacoblister/noisefloor/app/audiomodule/synth"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

// Processor is a synth processor block (with inputs and outputs displayed)
type Processor struct {
	ProcessorDefinition *synth.ProcessorDefinition

	editState editState
}

//MakeProcessor create an new Processor UI componenet
func MakeProcessor(processorDefinition *synth.ProcessorDefinition) *Processor {
	processor := Processor{ProcessorDefinition: processorDefinition}
	return &processor
}

// Render displays a processor
func (p *Processor) Render() vdom.Element {
	procWidth := 40
	procHeight := 80

	procName, procInputs, procOutputs := p.ProcessorDefinition.Processor.Definition()

	inConnectors := []vdom.Element{}
	for i := 0; i < len(procInputs); i++ {
		connector := vdom.MakeElement("rect",
			"id", procName+":inconn:"+strconv.Itoa(i),
			"x", p.ProcessorDefinition.X+0,
			"y", p.ProcessorDefinition.Y+(i+1)*8,
			"width", 4,
			"height", 4,
			"stroke", "black",
			"fill", "black",
		)
		inConnectors = append(inConnectors, connector)
	}

	outConnectors := []vdom.Element{}
	for i := 0; i < len(procOutputs); i++ {
		connector := vdom.MakeElement("rect",
			"id", procName+":inconn:"+strconv.Itoa(i),
			"x", p.ProcessorDefinition.X+procWidth-4,
			"y", p.ProcessorDefinition.Y+(i+1)*8,
			"width", 4,
			"height", 4,
			"stroke", "black",
			"fill", "black",
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
		),
		inConnectors,
		outConnectors,
	)

	return element
}
