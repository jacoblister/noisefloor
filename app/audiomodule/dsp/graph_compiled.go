package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

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

// compiledGraph the graph after compilation
type compiledGraph interface {
	Start(sampleRate int)
	Stop()
	Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event)
}

type graphOp struct {
	processor    Processor
	connectorIn  []*Connector
	connectorOut []*Connector
}

type graphExecutor struct {
	midiInput  *processorbuiltin.MIDIInput // 'special' MIDI input processor
	inputTerm  *processorbuiltin.Terminal  // 'special' Audio input terminal
	outputTerm *processorbuiltin.Terminal  // 'special' Audio output terminal
	ops        []graphOp                   // operations to perform, in order
}

func compileGraphExecutor(graph Graph) graphExecutor {
	graphExecutor := graphExecutor{}
	graphExecutor.midiInput = graph.ProcessorList[0].Processor.(*processorbuiltin.MIDIInput)
	graphExecutor.outputTerm = graph.ProcessorList[4].Processor.(*processorbuiltin.Terminal)

	nullConnector := Connector{}
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[0].Processor,
		[]*Connector{}, []*Connector{&graph.ConnectorList[0], &graph.ConnectorList[1], &graph.ConnectorList[2],
			&nullConnector, &nullConnector, &nullConnector, &nullConnector}})
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[1].Processor,
		[]*Connector{&graph.ConnectorList[0]}, []*Connector{&graph.ConnectorList[3]}})
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[2].Processor,
		[]*Connector{&graph.ConnectorList[1], &graph.ConnectorList[2]}, []*Connector{&graph.ConnectorList[4]}})
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[3].Processor,
		[]*Connector{&graph.ConnectorList[3], &graph.ConnectorList[4]}, []*Connector{&graph.ConnectorList[5]}})
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[4].Processor,
		[]*Connector{&graph.ConnectorList[5], &graph.ConnectorList[5]}, []*Connector{}})

	return graphExecutor
}

// compileProcessorGraph compiles a graph, and returns a function to run it
func compileProcessorGraph(graph Graph, target CompileTarget) compiledGraph {
	graphExecutor := compileGraphExecutor(graph)

	switch target {
	case CompileInterpreted:
		compiledGraph := interpretedEngine{graphExecutor: graphExecutor}
		return &compiledGraph

	}
	panic("unsupported target")
}
