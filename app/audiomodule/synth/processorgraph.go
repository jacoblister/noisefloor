package synth

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/synth/processor"
)

// ProcessorGraph is a graph of processors and connectors, plus exported parameter map
type ProcessorGraph struct {
	Name          string
	ProcessorList []ProcessorDefinition
	ConnectorList []Connector
}

// loadProcessorGraph loads a procesor graph from file
// just sets up a static graph for now
func loadProcessorGraph(filename string) ProcessorGraph {
	processorGraph := ProcessorGraph{}

	osc := processor.Oscillator{}
	osc.Freq = 440
	processorGraph.ProcessorList = append(processorGraph.ProcessorList,
		ProcessorDefinition{X: 100, Y: 100, Processor: &osc})
	env := processor.Envelope{}
	processorGraph.ProcessorList = append(processorGraph.ProcessorList,
		ProcessorDefinition{X: 100, Y: 200, Processor: &env})
	gain := processor.Gain{}
	processorGraph.ProcessorList = append(processorGraph.ProcessorList,
		ProcessorDefinition{X: 200, Y: 100, Processor: &gain})

	processorGraph.ConnectorList = append(processorGraph.ConnectorList,
		Connector{FromProcessor: &osc, FromPort: 0, ToProcessor: &gain, ToPort: 0})
	processorGraph.ConnectorList = append(processorGraph.ConnectorList,
		Connector{FromProcessor: &env, FromPort: 0, ToProcessor: &gain, ToPort: 1})

	return processorGraph
}
