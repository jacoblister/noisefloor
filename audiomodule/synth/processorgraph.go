package synth

import (
	"github.com/jacoblister/noisefloor/audiomodule/synth/processor"
)

// processorGraph is a graph of processors and connectors, plus exported parameter map
type processorGraph struct {
	name          string
	processorList []processorDefinition
	connectorList []connector
}

// loadProcessorGraph loads a procesor graph from file
// just sets up a static graph for now
func loadProcessorGraph(filename string) processorGraph {
	processorGraph := processorGraph{}

	osc := processor.Oscillator{}
	osc.Freq = 440
	processorGraph.processorList = append(processorGraph.processorList,
		processorDefinition{x: 100, y: 100, processor: &osc})
	env := processor.Envelope{}
	processorGraph.processorList = append(processorGraph.processorList,
		processorDefinition{x: 100, y: 200, processor: &env})
	gain := processor.Gain{}
	processorGraph.processorList = append(processorGraph.processorList,
		processorDefinition{x: 200, y: 100, processor: &gain})

	processorGraph.connectorList = append(processorGraph.connectorList,
		connector{fromProcessor: &osc, fromPort: 0, toProcessor: &gain, toPort: 0})
	processorGraph.connectorList = append(processorGraph.connectorList,
		connector{fromProcessor: &env, fromPort: 0, toProcessor: &gain, toPort: 1})

	return processorGraph
}
