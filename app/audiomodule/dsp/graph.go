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
		ProcessorDefinition{X: 80, Y: 80, Processor: &midiInput})
	osc := processor.Oscillator{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 240, Y: 80, Processor: &osc})
	env := processor.Envelope{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 240, Y: 240, Processor: &env})
	gain := processor.Gain{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 400, Y: 80, Processor: &gain})
	outputTerminal := processorbuiltin.Terminal{}
	outputTerminal.SetParameters(true, 2)
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 560, Y: 80, Processor: &outputTerminal})

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
		Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 0})
	// graph.ConnectorList = append(graph.ConnectorList,
	// 	Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 1})

	return graph
}
