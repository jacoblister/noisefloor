package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

// CompileTarget enumerated type
type CompileTarget int

// CompileTarget implementation
const (
	CompileInterpreted CompileTarget = iota
	CompileInterpretedSingleSample
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
	processor    processor.Processor
	connectorIn  []*processor.Connector
	connectorOut [][]*processor.Connector
}

type graphExecutor struct {
	midiInput  *processorbuiltin.MIDIInput // 'special' MIDI input processor
	inputTerm  *processorbuiltin.Terminal  // 'special' Audio input terminal
	outputTerm *processorbuiltin.Terminal  // 'special' Audio output terminal
	ops        []graphOp                   // operations to perform, in order
	connectors []processor.Connector       // all connectors in the graph
}

func compileGraphExecutor(graph Graph) graphExecutor {
	// init graphExecutor and provide default 'special' processors
	graphExecutor := graphExecutor{}
	graphExecutor.midiInput = &processorbuiltin.MIDIInput{}
	graphExecutor.outputTerm = &processorbuiltin.Terminal{}

	for i := 0; i < len(graph.Processors); i++ {
		// check for 'special' processors
		midiInput, ok := graph.Processors[i].Processor.(*processorbuiltin.MIDIInput)
		if ok {
			graphExecutor.midiInput = midiInput
		}
		terminal, ok := graph.Processors[i].Processor.(*processorbuiltin.Terminal)
		if ok {
			graphExecutor.outputTerm = terminal
		}

		graphExecutor.ops = append(graphExecutor.ops,
			graphOp{
				graph.Processors[i].Processor,
				graph.inputConnectorsForProcessor(graph.Processors[i].Processor),
				graph.outputConnectorsForProcessor(graph.Processors[i].Processor),
			},
		)
	}

	graphExecutor.connectors = graph.Connectors

	return graphExecutor
}

// compileProcessorGraph compiles a graph, and returns a function to run it
func compileProcessorGraph(graph Graph, target CompileTarget) compiledGraph {
	graphExecutor := compileGraphExecutor(graph)

	switch target {
	case CompileInterpreted:
		compiledGraph := interpretedEngine{graphExecutor: graphExecutor}
		return &compiledGraph
	case CompileInterpretedSingleSample:
		compiledGraph := interpretedEngineSingleSample{graphExecutor: graphExecutor}
		return &compiledGraph
	}
	panic("unsupported target")
}
