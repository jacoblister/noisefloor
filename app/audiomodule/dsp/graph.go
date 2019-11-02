package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
)

// Graph is a graph of processors and connectors, plus exported parameter map
type Graph struct {
	Name          string
	ProcessorList []ProcessorDefinition
	ConnectorList []Connector
}

func (g *Graph) connectorsForProcessor(processor Processor, isInput bool) []*Connector {
	_, procInputs, procOutputs := processor.Definition()
	connectorCount := 0
	if isInput {
		connectorCount = len(procInputs)
	} else {
		connectorCount = len(procOutputs)
	}
	result := make([]*Connector, connectorCount, connectorCount)
	for i := 0; i < len(result); i++ {
		result[i] = &Connector{}
	}

	for i := 0; i < len(g.ConnectorList); i++ {
		if isInput && g.ConnectorList[i].ToProcessor == processor {
			result[g.ConnectorList[i].ToPort] = &g.ConnectorList[i]
		}
		if !isInput && g.ConnectorList[i].FromProcessor == processor {
			result[g.ConnectorList[i].FromPort] = &g.ConnectorList[i]
		}
	}
	return result
}

// loadProcessorGraph loads a procesor graph from file
// just sets up a static graph for now
func loadProcessorGraph(filename string) Graph {
	graph := Graph{}

	midiInput := processorbuiltin.MIDIInput{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 16, Y: 16, Processor: &midiInput})
	osc := processor.Oscillator{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 120, Y: 16, Processor: &osc})
	env := processor.Envelope{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 120, Y: 72, Processor: &env})
	gain := processor.Gain{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 224, Y: 16, Processor: &gain})
	splitter := processor.Splitter{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 328, Y: 16, Processor: &splitter})
	outputTerminal := processorbuiltin.Terminal{}
	outputTerminal.SetParameters(true, 2)
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 432, Y: 16, Processor: &outputTerminal})
	scope := processor.Scope{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 432, Y: 96, Processor: &scope})

	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &midiInput, FromPort: 0, ToProcessor: &osc, ToPort: 0})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &midiInput, FromPort: 1, ToProcessor: &env, ToPort: 0})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &midiInput, FromPort: 2, ToProcessor: &env, ToPort: 1})

	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &osc, FromPort: 0, ToProcessor: &gain, ToPort: 0})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &env, FromPort: 0, ToProcessor: &gain, ToPort: 1})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &splitter, ToPort: 0})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &splitter, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 0})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &splitter, FromPort: 1, ToProcessor: &outputTerminal, ToPort: 1})
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &splitter, FromPort: 3, ToProcessor: &scope, ToPort: 0})

	return graph
}
