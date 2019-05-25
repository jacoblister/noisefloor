package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

// Graph is a graph of processors and connectors, plus exported parameter map
type Graph struct {
	Name          string
	ProcessorList []ProcessorDefinition
	ConnectorList []Connector
}

// loadProcessorGraph loads a procesor graph from file
// just sets up a static graph for now
func loadProcessorGraph(filename string) Graph {
	graph := Graph{}

	midiInput := processorbuiltin.MIDIInput{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 100, Y: 100, Processor: &midiInput})
	osc := processor.Oscillator{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 200, Y: 100, Processor: &osc})
	env := processor.Envelope{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 200, Y: 200, Processor: &env})
	gain := processor.Gain{}
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 300, Y: 100, Processor: &gain})
	outputTerminal := processorbuiltin.Terminal{}
	outputTerminal.SetParameters(true, 2)
	graph.ProcessorList = append(graph.ProcessorList,
		ProcessorDefinition{X: 400, Y: 100, Processor: &outputTerminal})

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
	graph.ConnectorList = append(graph.ConnectorList,
		Connector{FromProcessor: &gain, FromPort: 0, ToProcessor: &outputTerminal, ToPort: 1})

	return graph
}

// CompileTarget enumerated type
type CompileTarget int

// CompileTarget implementation
const (
	CompileInterpreted CompileTarget = iota
	CompileGolang
	CompileJavascript
	CompileWasm
	CompileCPP
)

// AudioProcessor is a frame based audio/midi processor
type AudioProcessor interface {
	Start(sampleRate int)
	Stop()
	Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event)
}

type graphOp struct {
	processor Processor
	inArgs    []int
	outArgs   []int
}

type graphExecutor struct {
	processor   []Processor
	inputNodes  []int
	outputNodes []int
	vars        []float32
	ops         []graphOp
}

// compileProcessorGraph compiles a graph, and returns a function to run it
func compileProcessorGraph(graph Graph, target CompileTarget) AudioProcessor {
	// dummy implementation for now

	switch target {
	case CompileGolang:
		return nil
	}
	panic("unsupported target")
}
