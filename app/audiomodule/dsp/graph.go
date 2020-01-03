package dsp

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbasic"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
)

// Graph is a graph of processors and connectors, plus exported parameter map
type Graph struct {
	Name       string
	Processors []ProcessorDefinition
	Connectors []Connector
}

func (g *Graph) inputConnectorsForProcessor(processor processor.Processor) []*Connector {
	_, procInputs, _, _ := processor.Definition()
	connectorCount := len(procInputs)
	result := make([]*Connector, connectorCount, connectorCount)
	for i := 0; i < len(result); i++ {
		result[i] = &Connector{}
	}

	for i := 0; i < len(g.Connectors); i++ {
		if g.Connectors[i].ToProcessor == processor {
			result[g.Connectors[i].ToPort] = &g.Connectors[i]
		}
	}
	return result
}

func (g *Graph) outputConnectorsForProcessor(processor processor.Processor) [][]*Connector {
	_, _, procOutputs, _ := processor.Definition()
	connectorCount := len(procOutputs)
	result := make([][]*Connector, connectorCount, connectorCount)
	for i := 0; i < len(result); i++ {
		result[i] = make([]*Connector, 0, 0)
	}

	for i := 0; i < len(g.Connectors); i++ {
		if g.Connectors[i].FromProcessor == processor {
			result[g.Connectors[i].FromPort] = append(result[g.Connectors[i].FromPort], &g.Connectors[i])
		}
	}
	return result
}

func (g *Graph) definitonForProcessor(processor processor.Processor) ProcessorDefinition {
	for i := 0; i < len(g.Processors); i++ {
		if g.Processors[i].Processor == processor {
			return g.Processors[i]
		}
	}
	panic("could not find processor definition")
}

func (g *Graph) getProcessorByName(name string) processor.Processor {
	for i := 0; i < len(g.Processors); i++ {
		if g.Processors[i].GetName() == name {
			return g.Processors[i].Processor
		}
	}
	panic("could not find processor definition")
}

func exampleGraph() Graph {
	graph := Graph{}

	midiInput := processorbuiltin.MIDIInput{}
	graph.Processors = append(graph.Processors,
		ProcessorDefinition{X: 16, Y: 16, Processor: &midiInput})
	osc := processorbasic.Oscillator{}
	processor.SetProcessorDefaults(&osc)
	graph.Processors = append(graph.Processors,
		ProcessorDefinition{X: 120, Y: 16, Processor: &osc})
	env := processorbasic.Envelope{}
	processor.SetProcessorDefaults(&env)
	graph.Processors = append(graph.Processors,
		ProcessorDefinition{X: 120, Y: 96, Processor: &env})
	gain := processorbasic.Gain{}
	processor.SetProcessorDefaults(&gain)
	graph.Processors = append(graph.Processors,
		ProcessorDefinition{X: 224, Y: 16, Processor: &gain})
	outputTerminal := processorbuiltin.Terminal{}
	outputTerminal.SetParameters(true, 2)
	graph.Processors = append(graph.Processors,
		ProcessorDefinition{X: 328, Y: 16, Processor: &outputTerminal})
	scope := processorbasic.Scope{Trigger: true, Skip: 4}
	graph.Processors = append(graph.Processors,
		ProcessorDefinition{X: 328, Y: 96, Processor: &scope})
	scope2 := processorbasic.Scope{Trigger: false, Skip: 200}
	graph.Processors = append(graph.Processors,
		ProcessorDefinition{X: 224, Y: 208, Name: "scope2", Processor: &scope2})

	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &midiInput, FromPort: 0, ToProcessor: &osc, ToPort: 0})
	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &midiInput, FromPort: 1, ToProcessor: &env, ToPort: 0})
	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &midiInput, FromPort: 2, ToProcessor: &env, ToPort: 1})

	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &osc, FromPort: 0, ToProcessor: &gain, ToPort: 0})
	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &env, FromPort: 0, ToProcessor: &gain, ToPort: 1})
	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 0})
	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 1})
	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &scope, ToPort: 0})
	graph.Connectors = append(graph.Connectors,
		Connector{FromProcessor: &env, FromPort: 0, ToProcessor: &scope2, ToPort: 0})

	return graph
}

// loadProcessorGraph loads a procesor graph from file
func loadProcessorGraph(reader io.Reader) (Graph, error) {
	byteValue, _ := ioutil.ReadAll(reader)

	var graph Graph
	err := xml.Unmarshal(byteValue, &graph)

	return graph, err

	// return exampleGraph(), nil
}

// saveProcessorGraph saves the graph to the provided writer
func saveProcessorGraph(graph Graph, writer io.Writer) {
	xml, _ := xml.MarshalIndent(&graph, "", "   ")
	writer.Write(xml)
	println(string(xml))
}
